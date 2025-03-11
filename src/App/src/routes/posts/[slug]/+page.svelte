<script>
	let { data } = $props();

	$inspect(data);
</script>

<div class="flex justify-center pt-4">
	<div class="flex flex-col content-center">
		<div class="flex gap-2">
			<div>{data.post.id}</div>
			<div>{data.post.user_id}</div>
			<div>{data.post.content}</div>
		</div>
	</div>
</div>
