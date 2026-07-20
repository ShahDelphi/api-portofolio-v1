# PowerShell Script to Test Portfolio API Local Connection

$BaseURL = "http://localhost:3000/api/v1"

Write-Host "Checking Server Health..." -ForegroundColor Cyan
try {
    $health = Invoke-RestMethod -Uri "http://localhost:3000/health" -Method Get
    Write-Host "Success: $($health.message)" -ForegroundColor Green
} catch {
    Write-Host "Failed to connect to local server. Make sure the Go backend is running on port 3000." -ForegroundColor Red
    Write-Host $_.Exception.Message -ForegroundColor Red
    exit
}

Write-Host "`nFetching public projects..." -ForegroundColor Cyan
try {
    $projects = Invoke-RestMethod -Uri "$BaseURL/projects" -Method Get
    Write-Host "Success: Found $($projects.data.Count) projects." -ForegroundColor Green
    foreach ($p in $projects.data) {
        Write-Host " - [$($p.order)] $($p.title)" -ForegroundColor Yellow
    }
} catch {
    Write-Host "Error fetching projects: $_" -ForegroundColor Red
}

Write-Host "`nFetching public skills..." -ForegroundColor Cyan
try {
    $skills = Invoke-RestMethod -Uri "$BaseURL/skills" -Method Get
    Write-Host "Success: Found $($skills.data.Count) skills." -ForegroundColor Green
} catch {
    Write-Host "Error fetching skills: $_" -ForegroundColor Red
}
