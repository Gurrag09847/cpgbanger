<script lang="ts">
	import { getSettings, updateSettings } from '../api';

	let { onSuccess }: { onSuccess: () => void } = $props();

	let enabled = $state(false);
	let intervalHours = $state(24);
	let intervalDays = $state(0);
	let loading = $state(true);
	let saving = $state(false);

	$effect(() => {
		loadSettings();
	});

	async function loadSettings() {
		try {
			const settings = await getSettings();
			enabled = settings['schedule_enabled'] === 'true';
			intervalHours = parseInt(settings['schedule_interval_hours'] || '24') || 24;
			intervalDays = parseInt(settings['schedule_interval_days'] || '0') || 0;
		} catch {
			// defaults
		} finally {
			loading = false;
		}
	}

	async function save() {
		saving = true;
		try {
			await updateSettings({
				schedule_enabled: enabled ? 'true' : 'false',
				schedule_interval_hours: String(intervalHours),
				schedule_interval_days: String(intervalDays)
			});
			onSuccess();
		} catch {
			// ignore
		} finally {
			saving = false;
		}
	}

	const inputClass =
		'border-border bg-background text-foreground w-20 rounded-md border px-3 py-2 text-sm outline-none focus:border-primary focus:ring-1 focus:ring-primary';
	const toggleClass =
		'bg-muted peer-focus:ring-primary peer-checked:bg-primary h-6 w-11 rounded-full after:bg-background after:border-border after:absolute after:start-[2px] after:top-[2px] after:h-5 after:w-5 after:rounded-full after:border after:transition-all peer-checked:after:translate-x-full peer-checked:after:border-white';
</script>

{#if loading}
	<p class="text-muted-foreground text-sm">Laddar schema...</p>
{:else}
	<div class="space-y-4">
		<label class="flex items-center gap-3">
			<div class="relative inline-block">
				<input type="checkbox" class="peer sr-only" bind:checked={enabled} onchange={save} />
				<div class={toggleClass}></div>
			</div>
			<span class="text-foreground text-sm font-medium">Aktivera schemalagd körning</span>
		</label>

		{#if enabled}
			<div class="flex items-center gap-3">
				<div>
					<label class="text-foreground mb-1 block text-xs font-medium" for="intervalHours"
						>Timmar</label
					>
					<input
						id="intervalHours"
						class={inputClass}
						type="number"
						min="0"
						bind:value={intervalHours}
					/>
				</div>
				<div>
					<label class="text-foreground mb-1 block text-xs font-medium" for="intervalDays"
						>Dagar</label
					>
					<input
						id="intervalDays"
						class={inputClass}
						type="number"
						min="0"
						bind:value={intervalDays}
					/>
				</div>
				<button
					class="bg-primary text-primary-foreground hover:bg-primary/90 mt-5 rounded-md px-3 py-2 text-xs font-medium disabled:opacity-50"
					onclick={save}
					disabled={saving}
				>
					{saving ? 'Sparar...' : 'Spara'}
				</button>
			</div>
			<p class="text-muted-foreground text-xs">
				{#if intervalDays > 0}
					Körs var {intervalDays} dag(ar)
				{:else if intervalHours > 0}
					Körs var {intervalHours} timme/timmar
				{/if}
				med aktiv fil och sparad konfiguration.
			</p>
		{/if}
	</div>
{/if}
