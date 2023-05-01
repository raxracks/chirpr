function like(id) {
	fetch("/api/v1/posts/like", {
    method: "PATCH",
    body: JSON.stringify({ id }),
    headers: {
        "Content-Type": "application/json"
    }
	});
}
