package main

const template = `
<!DOCTYPE html>
<html><head><meta charset="UTF-8">
<meta http-equiv="X-UA-Compatible" content="IE=edge">
<meta name="viewport" content="width=device-width, initial-scale=1.0"><link rel="stylesheet" href="./water.css"></head><body>
<div style="display: flex; align-items: center; align-self: center; flex-direction: row; flex-wrap: wrap; justify-content: flex-end;">
<button onclick="location.href='./index.html'">index</button>
<button onclick="location.href='./tags.html'">tags</button>
</div>%s</body></html>
<!-- Third Party Licenses : /THIRD_PARTY_LICENSES.md -->
`

const subTemplate = `
<!DOCTYPE html>
<html><head><meta charset="UTF-8">
<meta http-equiv="X-UA-Compatible" content="IE=edge">
<meta name="viewport" content="width=device-width, initial-scale=1.0"><link rel="stylesheet" href="../water.css"></head><body>
<div style="display: flex; align-items: center; align-self: center; flex-direction: row; flex-wrap: wrap; justify-content: flex-end;">
<button onclick="location.href='../index.html'">index</button>
<button onclick="location.href='../tags.html'">tags</button>
</div>%s</body></html>
<!-- Third Party Licenses : /THIRD_PARTY_LICENSES.md -->
`
