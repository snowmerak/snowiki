package main

var template = `
<!DOCTYPE html>
<html><head><link rel="stylesheet" href="./water.css"></head><body>
<div style="display: flex; align-items: center; align-self: center; flex-direction: row; flex-wrap: wrap; justify-content: flex-end;">
<button onclick="location.href='./search.html'">search</button>
<button onclick="location.href='./tags.html'">tags</button>
</div>%s</body></html>
`
