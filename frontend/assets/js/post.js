function post(body) {
	fetch("/api/v1/posts", {
    method: "POST",
    body: JSON.stringify({body, author: 1}),
    headers: {
        "Content-Type": "application/json"
    }
	});
}
