package crawler

import (
	"io"
	"net/http"
	"sync"
)

const UserAgent = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_6)" +
	" AppleWebKit/537.36 (KHTML, like Gecko)" +
	" Chrome/79.0.3945.130 Safari/537.36"

var httpClient = &http.Client{}

type Job struct {
	URL  string
	Link *NewsLink
}

func Crawler(index NewsIndex, useGoroutine bool) error {
	if !useGoroutine {
		for i := range index.Link {
			if err := fetchContent(index.Link[i].URI, &index.Link[i]); err != nil {
				return err
			}
		}
		return nil
	}

	// Worker pool로 병렬 처리(최대 10개 동시 연결)
	const numWorkers = 10
	jobs := make(chan Job, len(index.Link))
	errChan := make(chan error, 1)
	var wg sync.WaitGroup

	// Worker 시작
	for w := 0; w < numWorkers; w++ {
		wg.Add(1)
		go worker(jobs, errChan, &wg)
	}

	// Job 시작
	for i := range index.Link {
		jobs <- Job{
			URL:  index.Link[i].URI,
			Link: &index.Link[i],
		}
	}
	close(jobs)

	wg.Wait()
	close(errChan)

	return nil
}

func worker(jobs <-chan Job, errChan chan<- error, wg *sync.WaitGroup) {
	defer wg.Done()

	for job := range jobs {
		if err := fetchContent(job.URL, job.Link); err != nil {
			select {
			case errChan <- err:
			default:
			}
		}
	}
}

func fetchContent(url string, link *NewsLink) error {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	req.Header.Add("User-Agent", UserAgent)

	resp, err := httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	link.Content = string(body)
	return nil
}

// 하위 호환성을 위해 유지
func GetContent(url string, link *NewsLink, wg *sync.WaitGroup) error {
	defer wg.Done()
	return fetchContent(url, link)
}
