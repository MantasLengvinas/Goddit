export async function load({ params }) {
	const res = await fetch('http://localhost:8080/v1/posts');
	if (!res.ok) {
		throw new Error('Failed to fetch data!');
	}

	const posts = await res.json();
	return { posts };
}
