<script lang="ts">
	import '../app.css';

	let { children } = $props();
</script>

<main class="container mx-auto pt-12 pb-12">
	{@render children()}
</main>
