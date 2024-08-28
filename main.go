package main
import (
    "fmt"
    "github.com/gocolly/colly/v2"
    "os"
    "strings"
)
func main() {
    if len(os.Args) < 2 {
        fmt.Println("Please provide a letter to search for movies.")
        return
    }
    letter := strings.ToUpper(os.Args[1])
    c := colly.NewCollector()
    var movies []string
    c.OnHTML(".movie-card-title", func(e *colly.HTMLElement){
        title := e.Text
        if strings.HasPrefix(strings.ToUpper(title), letter){
            movies = append(movies, title)
        }
    })
    c.OnRequest(func(r *colly.Request) {
        fmt.Println("Visiting", r.URL)
    })
    c.OnError(func(r *colly.Response, err error) {
        fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "\n", err)
    })
    err := c.Visit("https://www.taste.io/")
    if err != nil {
        fmt.Println("Error visiting URL:", err)
        return
    }
    if len(movies) == 0{
        fmt.Println("No movies found starting with the letter", letter)
        return
    }

    fmt.Printf("Movies starting with the letter '%s':\n", letter)
    for _, movie := range movies {
        fmt.Println(movie)
    }
}
