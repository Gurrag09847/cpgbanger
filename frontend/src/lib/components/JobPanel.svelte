<script lang="ts">
	import { listJobs, cancelJob, getJobDownloadUrl, getJobSseUrl } from '../api';
	import type { JobRow, JobUpdate } from '../types';

	let { refresh, onUpdated }: { refresh: number; onUpdated: () => void } = $props();

	let jobs = $state<JobRow[]>([]);
	let loading = $state(true);
	let sseConnections = $state<Record<string, EventSource>>({});
	let runningJobs = $state<Record<string, JobUpdate>>({});

	async function load() {
		try {
			jobs = await listJobs();
			jobs.forEach((j) => {
				if (j.status === 'running') {
					connectSSE(j.id);
					runningJobs[j.id] = {
						status: 'running',
						progress: j.progress,
						processed: j.processed_rows,
						total: j.total_rows,
						error: ''
					};
				}
			});
		} catch {
			// ignore
		} finally {
			loading = false;
		}
	}

	function connectSSE(jobId: string) {
		if (sseConnections[jobId]) return;

		const es = new EventSource(getJobSseUrl(jobId));
		es.onmessage = (event) => {
			const update: JobUpdate = JSON.parse(event.data);
			runningJobs[jobId] = update;
			if (
				update.status === 'completed' ||
				update.status === 'cancelled' ||
				update.status === 'error'
			) {
				es.close();
				delete sseConnections[jobId];
				setTimeout(() => {
					delete runningJobs[jobId];
					onUpdated();
				}, 2000);
			}
		};
		es.onerror = () => {
			es.close();
			delete sseConnections[jobId];
		};
		sseConnections[jobId] = es;
	}

	async function handleCancel(jobId: string) {
		try {
			await cancelJob({ job_id: jobId });
		} catch {
			// ignore
		}
	}

	function statusColor(status: string): string {
		switch (status) {
			case 'completed':
				return 'text-success';
			case 'error':
				return 'text-destructive';
			case 'running':
				return 'text-primary';
			case 'cancelled':
				return 'text-warning';
			default:
				return 'text-muted-foreground';
		}
	}

	function statusLabel(status: string): string {
		switch (status) {
			case 'completed':
				return 'Klar';
			case 'error':
				return 'Fel';
			case 'running':
				return 'Körs';
			case 'cancelled':
				return 'Avbruten';
			default:
				return status;
		}
	}

	$effect(() => {
		refresh;
		load();
	});

	// Cleanup SSE connections
	$effect(() => {
		return () => {
			Object.values(sseConnections).forEach((es) => es.close());
		};
	});
</script>

{#if loading}
	<p class="text-muted-foreground text-sm">Laddar jobb...</p>
{:else if jobs.length === 0}
	<p class="text-muted-foreground text-sm">
		Inga jobb ännu. Ladda upp en fil och klicka "Kör nu" för att starta.
	</p>
{:else}
	<div class="overflow-x-auto">
		<table class="w-full text-sm">
			<thead>
				<tr class="border-border border-b text-left">
					<th class="text-muted-foreground pb-2 pr-4 font-medium">Fil</th>
					<th class="text-muted-foreground pb-2 pr-4 font-medium">Status</th>
					<th class="text-muted-foreground pb-2 pr-4 font-medium">Förlopp</th>
					<th class="text-muted-foreground pb-2 pr-4 font-medium">Tid</th>
					<th class="text-muted-foreground pb-2 font-medium">Åtgärder</th>
				</tr>
			</thead>
			<tbody>
				{#each jobs as job (job.id)}
					{@const live = runningJobs[job.id]}
					<tr class="border-border border-b">
						<td class="text-foreground max-w-[200px] truncate py-3 pr-4">{job.file_name}</td>
						<td class="py-3 pr-4">
							<span class={statusColor(job.status)}>
								{live ? statusLabel(live.status) : statusLabel(job.status)}
							</span>
						</td>
						<td class="py-3 pr-4 w-[200px]">
							<div class="flex items-center gap-2">
								<div class="bg-muted h-2 flex-1 overflow-hidden rounded-full">
									<div
										class="bg-primary h-full rounded-full transition-all duration-300"
										style="width: {Math.round(live?.progress ?? job.progress)}%"
									></div>
								</div>
								<span class="text-muted-foreground w-12 text-right text-xs">
									{live
										? `${live.processed}/${live.total}`
										: `${job.processed_rows}/${job.total_rows}`}
								</span>
							</div>
							{#if live?.error}
								<p class="text-destructive mt-1 text-xs">{live.error}</p>
							{/if}
							{#if job.error_message}
								<p class="text-destructive mt-1 text-xs">{job.error_message}</p>
							{/if}
						</td>
						<td class="text-muted-foreground py-3 pr-4 text-xs">{job.created_at}</td>
						<td class="py-3">
							<div class="flex gap-2">
								{#if job.status === 'running'}
									<button
										class="text-destructive hover:underline text-xs"
										onclick={() => handleCancel(job.id)}
									>
										Avbryt
									</button>
								{:else if job.status === 'completed'}
									<a
										href={getJobDownloadUrl(job.id)}
										class="text-primary hover:underline text-xs"
										download
									>
										Ladda ner
									</a>
								{/if}
							</div>
						</td>
					</tr>
				{/each}
			</tbody>
		</table>
	</div>
{/if}
