<script lang="ts">
	import { startFetch, getSettings, updateSettings } from '../api';
	import type { FetchRequest } from '../types';

	let {
		activeFileId,
		onSuccess
	}: { activeFileId: string | null; onSuccess: (jobId: string) => void } = $props();

	const companies = [
		{ value: 'illbruck', label: 'illbruck' },
		{ value: 'vandex', label: 'vandex' },
		{ value: 'nullifire', label: 'nullifire' },
		{ value: 'flowcrete', label: 'flowcrete' },
		{ value: 'matacryl', label: 'matacryl' },
		{ value: 'tremco', label: 'tremco' }
	];

	const docTypes = [
		{ value: '', label: '-- Välj --' },
		{ value: 'Produktdatablad', label: 'Produktdatablad' },
		{ value: 'Säkerhetsdatablad', label: 'Säkerhetsdatablad' },
		{ value: 'Prestandadeklaration', label: 'Prestandadeklaration' },
		{ value: 'Miljövarudeklaration (EPD)', label: 'Miljövarudeklaration (EPD)' },
		{ value: 'Certifikat', label: 'Certifikat' }
	];

	let form = $state({
		sheetName: '',
		sheetNamesStr: '',
		artikelnummerCol: '',
		pdfLinkCol: '',
		startNumber: 1,
		endNumber: 100,
		company: 'illbruck',
		documentType: '',
		useDocumentTypeColumn: false,
		documentTypeColumn: '',
		updateDate: false,
		dateColumn: '',
		concurrency: 5
	});

	let running = $state(false);
	let error = $state<string | null>(null);

	$effect(() => {
		loadSavedConfig();
	});

	async function loadSavedConfig() {
		try {
			const settings = await getSettings();
			const saved = settings['last_fetch_params'];
			if (saved) {
				const parsed = JSON.parse(saved);
				form = {
					...form,
					...parsed,
					concurrency: parsed.concurrency || 5,
					sheetNamesStr: parsed.sheetNames ? parsed.sheetNames.join(', ') : ''
				};
			}
		} catch {
			// use defaults
		}
	}

	async function handleSubmit(e: Event) {
		e.preventDefault();
		error = null;
		running = true;
		try {
			const sheetNames = form.sheetNamesStr
				? form.sheetNamesStr.split(',').map((s) => s.trim()).filter(Boolean)
				: undefined;

			const paramsJson = JSON.stringify({ ...form, sheetNames, sheetNamesStr: undefined });
			await updateSettings({ last_fetch_params: paramsJson });

			const req: FetchRequest = {
				sheetName: form.sheetName,
				artikelnummerCol: form.artikelnummerCol,
				pdfLinkCol: form.pdfLinkCol,
				startNumber: form.startNumber,
				endNumber: form.endNumber,
				company: form.company,
				documentType: form.documentType,
				use_document_type_column: form.useDocumentTypeColumn,
				document_type_column: form.documentTypeColumn,
				update_date: form.updateDate,
				date_column: form.dateColumn,
				concurrency: form.concurrency
			};
			if (sheetNames && sheetNames.length > 0) {
				req.sheetNames = sheetNames;
			}
			if (activeFileId) {
				req.upload_id = activeFileId;
			}
			const res = await startFetch(req);
			onSuccess(res.job_id);
		} catch (e: unknown) {
			error = e instanceof Error ? e.message : 'Kunde inte starta hämtning';
		} finally {
			running = false;
		}
	}

	const inputClass =
		'border-border bg-background placeholder:text-muted-foreground text-foreground w-full rounded-md border px-3 py-2 text-sm outline-none focus:border-primary focus:ring-1 focus:ring-primary';
	const labelClass = 'text-foreground mb-1 block text-sm font-medium';
	const checkboxClass =
		'border-border bg-background text-primary h-4 w-4 rounded border focus:ring-primary';
</script>

<form onsubmit={handleSubmit} class="space-y-4">
	<div class="grid gap-4 sm:grid-cols-2">
		<div>
			<label class={labelClass} for="sheetName">Bladnamn</label>
			<input
				id="sheetName"
				class={inputClass}
				type="text"
				bind:value={form.sheetName}
				placeholder="Blad1"
			/>
		</div>
		<div>
			<label class={labelClass} for="sheetNamesStr">Flera blad (kommaseparerat)</label>
			<input
				id="sheetNamesStr"
				class={inputClass}
				type="text"
				bind:value={form.sheetNamesStr}
				placeholder="t.ex. Blad1, Blad2"
			/>
		</div>
		<div>
			<label class={labelClass} for="company">Företag</label>
			<select id="company" class={inputClass} bind:value={form.company}>
				{#each companies as c}
					<option value={c.value}>{c.label}</option>
				{/each}
			</select>
		</div>
		<div>
			<label class={labelClass} for="artikelnummerCol">Artikelnummer-kolumn</label>
			<input
				id="artikelnummerCol"
				class={inputClass}
				type="text"
				bind:value={form.artikelnummerCol}
				placeholder="t.ex. Artikelnummer"
				required
			/>
		</div>
		<div>
			<label class={labelClass} for="pdfLinkCol">PDF-länk-kolumn</label>
			<input
				id="pdfLinkCol"
				class={inputClass}
				type="text"
				bind:value={form.pdfLinkCol}
				placeholder="t.ex. Filnamn eller webblänk"
			/>
		</div>
		<div>
			<label class={labelClass} for="startNumber">Startrad</label>
			<input
				id="startNumber"
				class={inputClass}
				type="number"
				min="1"
				bind:value={form.startNumber}
				required
			/>
		</div>
		<div>
			<label class={labelClass} for="endNumber">Slutrad</label>
			<input
				id="endNumber"
				class={inputClass}
				type="number"
				min="1"
				bind:value={form.endNumber}
				required
			/>
		</div>
		<div>
			<label class={labelClass} for="concurrency">Samtidiga anrop</label>
			<input
				id="concurrency"
				class={inputClass}
				type="number"
				min="1"
				max="20"
				bind:value={form.concurrency}
			/>
		</div>
	</div>

	<div class="border-border border-t pt-4">
		<label class="flex items-center gap-2 text-sm">
			<input type="checkbox" class={checkboxClass} bind:checked={form.useDocumentTypeColumn} />
			<span class="text-foreground">Använd dokumenttypskolumn från kalkylblad</span>
		</label>

		{#if form.useDocumentTypeColumn}
			<div class="mt-3">
				<label class={labelClass} for="docTypeCol">Dokumenttypskolumnens namn</label>
				<input
					id="docTypeCol"
					class={inputClass}
					type="text"
					bind:value={form.documentTypeColumn}
					placeholder="t.ex. Typ av dokument"
				/>
			</div>
		{:else}
			<div class="mt-3">
				<label class={labelClass} for="docType">Dokumenttyp</label>
				<select id="docType" class={inputClass} bind:value={form.documentType}>
					{#each docTypes as dt}
						<option value={dt.value}>{dt.label}</option>
					{/each}
				</select>
			</div>
		{/if}
	</div>

	<div class="border-border border-t pt-4">
		<label class="flex items-center gap-2 text-sm">
			<input type="checkbox" class={checkboxClass} bind:checked={form.updateDate} />
			<span class="text-foreground">Uppdatera datumkolumn på varje rad</span>
		</label>
		{#if form.updateDate}
			<div class="mt-3">
				<label class={labelClass} for="dateCol">Datumkolumnens namn</label>
				<input
					id="dateCol"
					class={inputClass}
					type="text"
					bind:value={form.dateColumn}
					placeholder="t.ex. Datum"
				/>
			</div>
		{/if}
	</div>

	{#if error}
		<p class="text-destructive text-sm">{error}</p>
	{/if}

	<button
		type="submit"
		class="bg-primary text-primary-foreground hover:bg-primary/90 inline-flex w-full items-center justify-center rounded-md px-4 py-2.5 text-sm font-medium transition-colors disabled:opacity-50"
		disabled={running || !activeFileId}
	>
		{#if running}
			Startar jobb...
		{:else}
			Kör nu
		{/if}
	</button>
	{#if !activeFileId}
		<p class="text-muted-foreground text-xs text-center">
			Välj en aktiv fil ovan för att aktivera kör-knappen
		</p>
	{/if}
</form>
