<script>
	let { data } = $props();

	// onstate to table data, for testing api calls
	$effect(() => {
		console.table(data);
	});
</script>

<div class="flex justify-center pt-4">
	<div class="flex flex-col content-center">
		{#each data.posts as post}
			<div class="flex gap-2">
				<div>{post.id}</div>
				<div>{post.user_id}</div>
				<div>{post.content}</div>
			</div>
		{/each}
	</div>
</div>
