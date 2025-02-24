<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <title>My Adventure</title>
  <script src="https://unpkg.com/@tailwindcss/browser@4"></script>
</head>
<body class="max-w-4xl mx-auto">
  <h1 class="bg-blue-200 h-10 mb-4 mx-auto pt-2 rounded-b-lg text-center">
    {{.Title}}
  </h1>
  <div>
    {{
      range .Story
    }}
      <p>
        {{.}}
      </p>
    {{end}}
  </div>
  <div class="flex flex-col gap-y-2">
    {{
      range .Options
    }}
    <a href="/{{.Arc}}" class="bg-blue-100 hover:bg-blue-200 mt-2 outline-1 outline-gray-400 p-2 rounded-sm transition-colors">
      {{.Text}}
    </a>
    {{end}}
  </div>
</body>
</html>