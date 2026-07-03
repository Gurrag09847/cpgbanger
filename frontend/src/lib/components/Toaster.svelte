<script lang="ts">
	import { toasts, removeToast } from '../toast.svelte';
</script>

<div class="pointer-events-none fixed right-4 top-4 z-50 flex flex-col gap-2">
	{#each $toasts as toast (toast.id)}
		<button
			class="pointer-events-auto border-border bg-card text-foreground max-w-sm rounded-lg border px-4 py-3 text-left text-sm shadow-lg {toast.variant ===
			'error'
				? 'border-red-500'
				: ''} {toast.variant === 'success' ? 'border-green-500' : ''}"
			onclick={() => removeToast(toast.id)}
		>
			{toast.message}
		</button>
	{/each}
</div>
