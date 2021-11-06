$mainGitHubUrl="https://github.com/kat-co/concurrency-in-go-src"
Write-Host "PROJECT helpers loaded"

function Open-Github {
    $browserPath = GET-DefaultBrowserPath
    
    Start-Process $browserPath -ArgumentList $mainGitHubUrl
}