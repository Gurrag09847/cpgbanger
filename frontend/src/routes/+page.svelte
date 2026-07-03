<script lang="ts">
	import FileUpload from '../lib/components/FileUpload.svelte';
	import FileSelector from '../lib/components/FileSelector.svelte';
	import ConfigForm from '../lib/components/ConfigForm.svelte';
	import SchedulePanel from '../lib/components/SchedulePanel.svelte';
	import JobPanel from '../lib/components/JobPanel.svelte';
	import Toaster from '../lib/components/Toaster.svelte';
	import { addToast } from '../lib/toast.svelte';

	let activeFileId = $state<string | null>(null);
	let activeFileName = $state<string | null>(null);
	let refreshJobs = $state(0);
</script>

<div class="space-y-8">
	<div class="grid gap-8 lg:grid-cols-2">
		<section class="space-y-6">
			<div class="border-border bg-card rounded-lg border p-6">
				<h2 class="text-foreground mb-4 text-lg font-semibold">1. Ladda upp Excel-fil</h2>
				<FileUpload
					onUploaded={(id, name) => {
						activeFileId = id;
						activeFileName = name;
						addToast(`Uppladdad: ${name}`);
					}}
				/>
			</div>

			<div class="border-border bg-card rounded-lg border p-6">
				<h2 class="text-foreground mb-4 text-lg font-semibold">2. Välj fil</h2>
				<FileSelector bind:activeFileId bind:activeFileName onRefresh={() => refreshJobs++} />
			</div>
		</section>

		<section class="space-y-6">
			<div class="border-border bg-card rounded-lg border p-6">
				<h2 class="text-foreground mb-4 text-lg font-semibold">3. Hämtningsinställningar</h2>
				<ConfigForm
					{activeFileId}
					onSuccess={(jobId) => {
						refreshJobs++;
						addToast(`Jobb startat: ${jobId.slice(0, 8)}...`);
					}}
				/>
			</div>

			<div class="border-border bg-card rounded-lg border p-6">
				<h2 class="text-foreground mb-4 text-lg font-semibold">4. Schemalägg</h2>
				<SchedulePanel onSuccess={() => addToast('Schema uppdaterat')} />
			</div>
		</section>
	</div>

	<div class="border-border bg-card rounded-lg border p-6">
		<h2 class="text-foreground mb-4 text-lg font-semibold">Jobbhistorik</h2>
		<JobPanel refresh={refreshJobs} onUpdated={() => {}} />
	</div>
</div>

<Toaster />
