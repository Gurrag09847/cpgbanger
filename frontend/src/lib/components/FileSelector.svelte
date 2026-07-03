<script lang="ts">
	import {
		listUploads,
		deleteUpload,
		getSettings,
		updateSettings,
		getUploadDownloadUrl
	} from '../api';
	import type { UploadRow } from '../types';

	let {
		activeFileId = $bindable(null),
		activeFileName = $bindable(null),
		onRefresh
	}: {
		activeFileId: string | null;
		activeFileName: string | null;
		onRefresh: () => void;
	} = $props();

	let uploads = $state<UploadRow[]>([]);
	let loading = $state(true);

	async function load() {
		loading = true;
		try {
			uploads = await listUploads();
			const settings = await getSettings();
			if (settings['active_file_id']) {
				const found = uploads.find((u) => u.id === settings['active_file_id']);
				if (found) {
					activeFileId = found.id;
					activeFileName = found.original_name;
				}
			}
		} catch {
			// ignore
		} finally {
			loading = false;
		}
	}

	async function selectFile(id: string) {
		await updateSettings({ active_file_id: id });
		const found = uploads.find((u) => u.id === id);
		if (found) {
			activeFileId = found.id;
			activeFileName = found.original_name;
		}
	}

	async function removeFile(id: string) {
		await deleteUpload(id);
		if (activeFileId === id) {
			await updateSettings({ active_file_id: '' });
			activeFileId = null;
			activeFileName = null;
		}
		uploads = uploads.filter((u) => u.id !== id);
		onRefresh();
	}

	$effect(() => {
		load();
	});
</script>

{#if loading}
	<p class="text-muted-foreground text-sm">Laddar filer...</p>
{:else if uploads.length === 0}
	<p class="text-muted-foreground text-sm">Inga filer uppladdade ännu. Ladda upp en ovan.</p>
{:else}
	<ul class="divide-border divide-y rounded-lg border">
		{#each uploads as upload (upload.id)}
			<li
				class="flex items-center justify-between px-4 py-3 text-sm transition-colors {upload.id ===
				activeFileId
					? 'bg-primary/10'
					: ''}"
			>
				<div class="min-w-0 flex-1">
					<p class="text-foreground truncate font-medium">{upload.original_name}</p>
					<p class="text-muted-foreground text-xs">{upload.uploaded_at}</p>
				</div>
				<div class="ml-3 flex items-center gap-2">
					<a
						href={getUploadDownloadUrl(upload.id)}
						class="text-primary hover:underline text-xs"
						download
					>
						Ladda ner
					</a>
					{#if upload.id === activeFileId}
						<span class="bg-success/20 text-success rounded px-2 py-0.5 text-xs font-medium"
							>Aktiv</span
						>
					{:else}
						<button
							class="bg-primary text-primary-foreground hover:bg-primary/90 rounded px-2 py-0.5 text-xs font-medium"
							onclick={() => selectFile(upload.id)}
						>
							Välj
						</button>
					{/if}
					<button
						class="text-destructive hover:underline text-xs"
						onclick={() => removeFile(upload.id)}
					>
						Ta bort
					</button>
				</div>
			</li>
		{/each}
	</ul>
{/if}
