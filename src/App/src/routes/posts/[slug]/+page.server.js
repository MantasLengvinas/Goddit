export async function load({ params }) {
	const res = await fetch(`http://localhost:8080/v1/posts/${params.slug}`);
	if (!res.ok) {
		throw new Error('Failed to fetch data!');
	}

	const post = await res.json();
	return { post };
}
