<script lang="ts">
	import { uploadFile } from '../api';

	let { onUploaded }: { onUploaded: (id: string, name: string) => void } = $props();

	let dragging = $state(false);
	let uploading = $state(false);
	let error = $state<string | null>(null);

	async function handleFiles(files: FileList | null) {
		if (!files || files.length === 0) return;
		const file = files[0];
		if (!file.name.endsWith('.xlsx')) {
			error = 'Endast .xlsx-filer accepteras';
			return;
		}
		error = null;
		uploading = true;
		try {
			const res = await uploadFile(file);
			onUploaded(res.file_id, res.file_name);
		} catch (e: unknown) {
			error = e instanceof Error ? e.message : 'Uppladdning misslyckades';
		} finally {
			uploading = false;
		}
	}

	function onDrop(e: DragEvent) {
		e.preventDefault();
		dragging = false;
		handleFiles(e.dataTransfer?.files ?? null);
	}
</script>

<div
	class="border-border hover:border-primary/50 bg-muted cursor-pointer rounded-lg border-2 border-dashed p-8 text-center transition-colors {dragging
		? 'border-primary'
		: ''} {uploading ? 'opacity-50' : ''}"
	ondragover={(e) => {
		e.preventDefault();
		dragging = true;
	}}
	ondragleave={() => (dragging = false)}
	ondrop={onDrop}
	onclick={() => document.getElementById('file-input')?.click()}
	role="button"
	tabindex="0"
	onkeydown={(e) => {
		if (e.key === 'Enter') document.getElementById('file-input')?.click();
	}}
>
	{#if uploading}
		<p class="text-foreground">Laddar upp...</p>
	{:else}
		<svg
			class="text-muted-foreground mx-auto mb-3 h-10 w-10"
			fill="none"
			stroke="currentColor"
			viewBox="0 0 24 24"
		>
			<path
				stroke-linecap="round"
				stroke-linejoin="round"
				stroke-width="1.5"
				d="M7 16a4 4 0 01-.88-7.903A5 5 0 1115.9 6L16 6a5 5 0 011 9.9M15 13l-3-3m0 0l-3 3m3-3v12"
			/>
		</svg>
		<p class="text-foreground font-medium">Släpp din .xlsx-fil här</p>
		<p class="text-muted-foreground mt-1 text-sm">eller klicka för att bläddra</p>
	{/if}
</div>
<input
	id="file-input"
	type="file"
	accept=".xlsx"
	class="hidden"
	onchange={(e) => handleFiles((e.target as HTMLInputElement).files)}
/>
{#if error}
	<p class="text-destructive mt-2 text-sm">{error}</p>
{/if}
