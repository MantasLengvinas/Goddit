export async function load({ fetch }) {
	const res = await fetch('http://localhost:8080/posts');
	if (!res.ok) {
		throw new Error('Failed to fetch data!');
	}

	const posts = await res.json();
	return { posts };
}
