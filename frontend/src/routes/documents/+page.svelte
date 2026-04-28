<!-- <script lang="ts">
	import * as Form from '$lib/components/ui/form/index';
	import { Input } from '$lib/components/ui/input/index';
	import * as Select from '$lib/components/ui/select/index';
	import { formSchema, type FormSchema } from './schema';
	import { type SuperValidated, type Infer, superForm } from 'sveltekit-superforms';
	import { zod4Client } from 'sveltekit-superforms/adapters';
	import Loader from '@lucide/svelte/icons/loader';
	import { FetchDocuments, CancelFetch, SelectExcelFile } from '../../../wailsjs/go/main/App';
	import * as rt from '../../../wailsjs/runtime';
	import { Progress } from '$lib/components/ui/progress';
	import { onMount } from 'svelte';
	import { Button } from '$lib/components/ui/button';

	let { data }: { data: { form: SuperValidated<Infer<FormSchema>> } } = $props();

	const form = superForm(data.form, {
		validators: zod4Client(formSchema)
	});

	const { form: formData, enhance, validateForm } = form;

	let loading = $state(false);
	let progress = $state(0);
	let cancelled = $state(false);
	let selectedFilePath = $state('');

	onMount(() => {
		rt.EventsOn('progress_update', (val) => {
			progress = val;
		});
		rt.EventsOn('progress_cancelled', () => {
			cancelled = true;
			loading = false;
			progress = 0;
		});
	});

	const selectFile = async () => {
		const fileResult = await SelectExcelFile();

		if (fileResult) {
			selectedFilePath = fileResult;
			$formData.fileName = fileResult;
		}
	};

	const onSubmit = async (e: SubmitEvent) => {
		e.preventDefault();
		e.stopPropagation();
		progress = 0;
		cancelled = false;
		loading = true;
		const validatedForm = await validateForm({ focusOnError: true, update: true });

		if (!validatedForm.valid) {
			loading = false;
			return;
		}

		try {
			const res = await FetchDocuments({
				startNumber: validatedForm.data.startNumber,
				endNumber: validatedForm.data.endNumber,
				file: selectedFilePath,
				sheetName: validatedForm.data.sheetName,
				documentType: validatedForm.data.documentType
			});
		} catch (err) {
			console.log(err);
		} finally {
			loading = false;
		}
	};
</script>

<form onsubmit={onSubmit} method="POST" use:enhance>
	<div class="grid grid-cols-1 gap-3 sm:grid-cols-2">
		<Form.Field {form} name="startNumber">
			<Form.Control>
				{#snippet children({ props })}
					<Form.Label>Startnummer</Form.Label>
					<Input {...props} bind:value={$formData.startNumber} type="number" min={0} />
				{/snippet}
			</Form.Control>
			<Form.Description>Detta är numret du vill börja lägga till dokument på.</Form.Description>
			<Form.FieldErrors />
		</Form.Field>

		<Form.Field {form} name="endNumber">
			<Form.Control>
				{#snippet children({ props })}
					<Form.Label>Slutnummer</Form.Label>
					<Input {...props} bind:value={$formData.endNumber} type="number" min={0} />
				{/snippet}
			</Form.Control>
			<Form.Description>Detta är numret på sista raden du vill ha ett dokument på.</Form.Description
			>
			<Form.FieldErrors />
		</Form.Field>
	</div>

	<div class="grid grid-cols-1 gap-3 sm:grid-cols-2">
		<Form.Field {form} name="documentType">
			<Form.Control>
				{#snippet children({ props })}
					<Form.Label>Dokumenttyp</Form.Label>
					<Select.Root type="single" bind:value={$formData.documentType} name={props.name}>
						<Select.Trigger {...props} class="w-full">
							{$formData.documentType ? $formData.documentType : 'Välj en dokumenttyp'}
						</Select.Trigger>
						<Select.Content class="w-full">
							<Select.Item value="Produktdatablad" label="Produktdatablad" />
							<Select.Item value="Säkerhetsdatablad" label="Säkerhetsdatablad" />
							<Select.Item value="Prestandadeklaration" label="Prestandadeklaration" />
						</Select.Content>
					</Select.Root>
				{/snippet}
			</Form.Control>
			<Form.Description>Detta är typen av dokument du vill lägga till.</Form.Description>
			<Form.FieldErrors />
		</Form.Field>

		<Form.Field {form} name="sheetName">
			<Form.Control>
				{#snippet children({ props })}
					<Form.Label>Namn på excel-ark</Form.Label>
					<Select.Root type="single" bind:value={$formData.sheetName} name={props.name}>
						<Select.Trigger {...props} class="w-full">
							{$formData.sheetName ? $formData.sheetName : 'Välj ett excel-ark'}
						</Select.Trigger>
						<Select.Content class="w-full">
							<Select.Item value="Dokument" label="Dokument" />
							<Select.Item value="illbruck" label="illbruck" />
						</Select.Content>
					</Select.Root>
				{/snippet}
			</Form.Control>
			<Form.Description>Detta är namnet på excel-arket du vill använda.</Form.Description>
			<Form.FieldErrors />
		</Form.Field>
	</div>

	<div>
		<Form.Field {form} name="fileName">
			<Form.Control>
				<Form.Label>Excel-fil</Form.Label>
				<div class="flex items-center gap-2">
					<Button type="button" onclick={selectFile}>Välj fil</Button>
					<span class="truncate text-sm italic">{selectedFilePath}</span>
				</div>
			</Form.Control>
			<Form.Description>Välj din Excel-fil.</Form.Description>
			<Form.FieldErrors />
		</Form.Field>
	</div>

	<div class="flex justify-end">
		<Form.Button disabled={loading} class="mt-8 w-full sm:w-auto">
			{#if loading}
				<Loader class="h-4 w-4 animate-spin" /> Arbetar...
			{:else}
				Börja leta efter dokument
			{/if}
		</Form.Button>
	</div>

	{#if cancelled}
		<p class="py-14 text-center text-lg font-medium">Operationen avbröts</p>
	{:else if progress}
		<div class="py-8">
			<Progress value={progress} />
			<div class="flex justify-between py-2">
				{#if progress < 100}
					<Button type="button" onclick={() => CancelFetch()}>Avbryt</Button>
				{/if}
				<p class="text-sm font-medium">{progress.toFixed(2)}%</p>
			</div>
		</div>
	{/if}
</form> -->

<!-- <script lang="ts">
	import * as Form from '$lib/components/ui/form/index';
	import { Input } from '$lib/components/ui/input/index';
	import * as Select from '$lib/components/ui/select/index';
	import { Checkbox } from '$lib/components/ui/checkbox';
	import { formSchema, type FormSchema } from './schema';
	import { type SuperValidated, type Infer, superForm } from 'sveltekit-superforms';
	import { zod4Client } from 'sveltekit-superforms/adapters';
	import Loader from '@lucide/svelte/icons/loader';
	import { FetchDocuments, CancelFetch, SelectExcelFile } from '../../../wailsjs/go/main/App';
	import * as rt from '../../../wailsjs/runtime';
	import { Progress } from '$lib/components/ui/progress';
	import { onMount } from 'svelte';
	import { Button } from '$lib/components/ui/button';

	let { data }: { data: { form: SuperValidated<Infer<FormSchema>> } } = $props();

	const form = superForm(data.form, {
		validators: zod4Client(formSchema)
	});

	const { form: formData, enhance, validateForm } = form;

	let loading = $state(false);
	let progress = $state(0);
	let cancelled = $state(false);
	let selectedFilePath = $state('');

	onMount(() => {
		rt.EventsOn('progress_update', (val) => {
			progress = val;
		});
		rt.EventsOn('progress_cancelled', () => {
			cancelled = true;
			loading = false;
			progress = 0;
		});
	});

	const selectFile = async () => {
		const fileResult = await SelectExcelFile();

		if (fileResult) {
			selectedFilePath = fileResult;
			$formData.fileName = fileResult;
		}
	};

	const onSubmit = async (e: SubmitEvent) => {
		e.preventDefault();
		e.stopPropagation();
		progress = 0;
		cancelled = false;
		loading = true;
		const validatedForm = await validateForm({ focusOnError: true, update: true });

		if (!validatedForm.valid) {
			loading = false;
			return;
		}

		try {
			const res = await FetchDocuments({
				startNumber: validatedForm.data.startNumber,
				endNumber: validatedForm.data.endNumber,
				file: selectedFilePath,
				sheetName: validatedForm.data.sheetName,
				documentType: validatedForm.data.documentType,
				artikelnummerCol: validatedForm.data.articleNumberColumn,
				pdfLinkCol: validatedForm.data.pdfLinkColumn,
				company: validatedForm.data.company,
			update_date: validatedForm.data.updateDate,
			date_column: validatedForm.data.dateColumn,
			document_type_column: validatedForm.
			});
		} catch (err) {
			console.log(err);
		} finally {
			loading = false;
		}
	};
</script>

<form onsubmit={onSubmit} method="POST" use:enhance>
	<div class="grid grid-cols-1 gap-3 sm:grid-cols-2">
		<Form.Field {form} name="startNumber">
			<Form.Control>
				{#snippet children({ props })}
					<Form.Label>Startnummer</Form.Label>
					<Input {...props} bind:value={$formData.startNumber} type="number" min={0} />
				{/snippet}
			</Form.Control>
			<Form.Description>Detta är numret du vill börja lägga till dokument på.</Form.Description>
			<Form.FieldErrors />
		</Form.Field>

		<Form.Field {form} name="endNumber">
			<Form.Control>
				{#snippet children({ props })}
					<Form.Label>Slutnummer</Form.Label>
					<Input {...props} bind:value={$formData.endNumber} type="number" min={0} />
				{/snippet}
			</Form.Control>
			<Form.Description>Detta är numret på sista raden du vill ha ett dokument på.</Form.Description>
			<Form.FieldErrors />
		</Form.Field>
	</div>

	<div class="grid grid-cols-1 gap-3 sm:grid-cols-2">
		<Form.Field {form} name="documentType">
			<Form.Control>
				{#snippet children({ props })}
					<Form.Label>Dokumenttyp</Form.Label>
					<Select.Root type="single" bind:value={$formData.documentType} name={props.name}>
						<Select.Trigger {...props} class="w-full">
							{$formData.documentType ? $formData.documentType : 'Välj en dokumenttyp'}
						</Select.Trigger>
						<Select.Content class="w-full">
							<Select.Item value="Produktdatablad" label="Produktdatablad" />
							<Select.Item value="Säkerhetsdatablad" label="Säkerhetsdatablad" />
							<Select.Item value="Prestandadeklaration" label="Prestandadeklaration" />
						</Select.Content>
					</Select.Root>
				{/snippet}
			</Form.Control>
			<Form.Description>Detta är typen av dokument du vill lägga till.</Form.Description>
			<Form.FieldErrors />
		</Form.Field>

		<Form.Field {form} name="sheetName">
			<Form.Control>
				{#snippet children({ props })}
					<Form.Label>Namn på excel-ark</Form.Label>
					<Input {...props} bind:value={$formData.sheetName} type="text"  />

				{/snippet}
			</Form.Control>
			<Form.Description>Detta är namnet på excel-arket du vill använda.</Form.Description>
			<Form.FieldErrors />
		</Form.Field>

		<Form.Field {form} name="company">
			<Form.Control>
				{#snippet children({ props })}
					<Form.Label>Välj ett märke att hämta information ifrån</Form.Label>
					<Select.Root type="single" bind:value={$formData.company} name={props.name}>
						<Select.Trigger {...props} class="w-full">
							{$formData.company ? $formData.company.at(0)?.toUpperCase() + $formData.company.slice(1, $formData.company.length) : 'Välj ett märke'}
						</Select.Trigger>
						<Select.Content class="w-full">
							<Select.Item value="illbruck" label="Illbruck" />
							<Select.Item value="nullifire" label="Nullifire" />
							<Select.Item value="matacryl" label="Matacryl" />
							<Select.Item value="flowcrete" label="Flowcrete" />
							<Select.Item value="vandex" label="Vandex" />
						</Select.Content>
					</Select.Root>
				{/snippet}
			</Form.Control>
			<Form.Description>Detta är typen av dokument du vill lägga till.</Form.Description>
			<Form.FieldErrors />
		</Form.Field>
	</div>

	<div>
		<Form.Field {form} name="fileName">
			<Form.Control>
				<Form.Label>Excel-fil</Form.Label>
				<div class="flex items-center gap-2">
					<Button type="button" onclick={selectFile}>Välj fil</Button>
					<span class="truncate text-sm italic">{selectedFilePath}</span>
				</div>
			</Form.Control>
			<Form.Description>Välj din Excel-fil.</Form.Description>
			<Form.FieldErrors />
		</Form.Field>
	</div>

	<div class="mt-6 space-y-4 rounded-lg border p-4">
		<h3 class="text-lg font-semibold">Kolumnkonfiguration</h3>

		<div class="grid grid-cols-1 gap-3 sm:grid-cols-2">
			<Form.Field {form} name="articleNumberColumn">
				<Form.Control>
					{#snippet children({ props })}
						<Form.Label>Artikelnummerkolumn</Form.Label>
						<Input {...props} bind:value={$formData.articleNumberColumn} type="text" />
					{/snippet}
				</Form.Control>
				<Form.Description>Namnet på kolumnen med artikelnummer.</Form.Description>
				<Form.FieldErrors />
			</Form.Field>

			<Form.Field {form} name="pdfLinkColumn">
			<Form.Control>
				{#snippet children({ props })}
					<Form.Label>PDF-länkkolumn</Form.Label>
					<Input {...props} bind:value={$formData.pdfLinkColumn} type="text" />
				{/snippet}
			</Form.Control>
			<Form.Description>Namnet på kolumnen där PDF-länkar ska sparas.</Form.Description>
			<Form.FieldErrors />
		</Form.Field>
		</div>



		<Form.Field {form} name="updateDate" class="flex flex-row items-start space-x-3 space-y-0 rounded-md border p-4">
			<Form.Control>
				{#snippet children({ props })}
					<Checkbox {...props} bind:checked={$formData.updateDate} />
					<div class="space-y-1 leading-none">
						<Form.Label>Uppdatera datum?</Form.Label>
						<Form.Description>
							Om denna är aktiverad kommer datumet att uppdateras, men du måste välja vilken kolumn.
						</Form.Description>
					</div>
				{/snippet}
			</Form.Control>
		</Form.Field>

		{#if $formData.updateDate}
			<Form.Field {form} name="dateColumn">
				<Form.Control>
					{#snippet children({ props })}
						<Form.Label>Datumkolumn</Form.Label>
						<Input {...props} bind:value={$formData.dateColumn} type="text" />
					{/snippet}
				</Form.Control>
				<Form.Description>Namnet på kolumnen för datum.</Form.Description>
				<Form.FieldErrors />
			</Form.Field>
		{/if}
	</div>

	<div class="flex justify-end">
		<Form.Button disabled={loading} class="mt-8 w-full sm:w-auto">
			{#if loading}
				<Loader class="h-4 w-4 animate-spin" /> Arbetar...
			{:else}
				Börja leta efter dokument
			{/if}
		</Form.Button>
	</div>

	{#if cancelled}
		<p class="py-14 text-center text-lg font-medium">Operationen avbröts</p>
	{:else if progress}
		<div class="py-8">
			<Progress value={progress} />
			<div class="flex justify-between py-2">
				{#if progress < 100}
					<Button type="button" onclick={() => CancelFetch()}>Avbryt</Button>
				{/if}
				<p class="text-sm font-medium">{progress.toFixed(2)}%</p>
			</div>
		</div>
	{/if}
</form> -->

<script lang="ts">
	import * as Form from '$lib/components/ui/form/index';
	import { Input } from '$lib/components/ui/input/index';
	import * as Select from '$lib/components/ui/select/index';
	import { Checkbox } from '$lib/components/ui/checkbox';
	import { formSchema, type FormSchema } from './schema';
	import { type SuperValidated, type Infer, superForm } from 'sveltekit-superforms';
	import { zod4Client } from 'sveltekit-superforms/adapters';
	import Loader from '@lucide/svelte/icons/loader';
	import { FetchDocuments, CancelFetch, SelectExcelFile } from '../../../wailsjs/go/main/App';
	import * as rt from '../../../wailsjs/runtime';
	import { Progress } from '$lib/components/ui/progress';
	import { onMount } from 'svelte';
	import { Button } from '$lib/components/ui/button';

	let { data }: { data: { form: SuperValidated<Infer<FormSchema>> } } = $props();

	const form = superForm(data.form, {
		validators: zod4Client(formSchema)
	});

	const { form: formData, enhance, validateForm } = form;

	let loading = $state(false);
	let progress = $state(0);
	let cancelled = $state(false);
	let selectedFilePath = $state('');

	onMount(() => {
		rt.EventsOn('progress_update', (val) => {
			progress = val;
		});
		rt.EventsOn('progress_cancelled', () => {
			cancelled = true;
			loading = false;
			progress = 0;
		});
	});

	const selectFile = async () => {
		const fileResult = await SelectExcelFile();

		if (fileResult) {
			selectedFilePath = fileResult;
			$formData.fileName = fileResult;
		}
	};

	const onSubmit = async (e: SubmitEvent) => {
		e.preventDefault();
		progress = 0;
		cancelled = false;
		loading = true;

		const validatedForm = await validateForm({ focusOnError: true, update: true });

		if (!validatedForm.valid) {
			loading = false;
			return;
		}

		try {
			await FetchDocuments({
				startNumber: validatedForm.data.startNumber,
				endNumber: validatedForm.data.endNumber,
				file: selectedFilePath,
				sheetName: validatedForm.data.sheetName,
				documentType: validatedForm.data.documentType,
				artikelnummerCol: validatedForm.data.articleNumberColumn,
				pdfLinkCol: validatedForm.data.pdfLinkColumn,
				company: validatedForm.data.company,
				update_date: validatedForm.data.updateDate,
				date_column: validatedForm.data.dateColumn,
				use_document_type_column: validatedForm.data.useDocumentTypeColumn,
				document_type_column: validatedForm.data.documentTypeColumn
			});
		} catch (err) {
			console.log(err);
		} finally {
			loading = false;
		}
	};
</script>

<form onsubmit={onSubmit} method="POST" use:enhance>
	<!-- Start/End Number -->
	<div class="grid grid-cols-1 gap-3 sm:grid-cols-2">
		<Form.Field {form} name="startNumber">
			<Form.Control>
				{#snippet children({ props })}
					<Form.Label>Startnummer</Form.Label>
					<Input {...props} bind:value={$formData.startNumber} type="number" min={0} />
				{/snippet}
			</Form.Control>
			<Form.Description>Numret du vill börja på.</Form.Description>
			<Form.FieldErrors />
		</Form.Field>

		<Form.Field {form} name="endNumber">
			<Form.Control>
				{#snippet children({ props })}
					<Form.Label>Slutnummer</Form.Label>
					<Input {...props} bind:value={$formData.endNumber} type="number" min={0} />
				{/snippet}
			</Form.Control>
			<Form.Description>Numret du vill sluta på.</Form.Description>
			<Form.FieldErrors />
		</Form.Field>
	</div>

	<!-- Document Type & Sheet Name -->
	<div class="grid grid-cols-1 gap-3 sm:grid-cols-2">
		<Form.Field {form} name="useDocumentTypeColumn" class="flex gap-3 rounded-md border p-4">
			<Form.Control>
				{#snippet children({ props })}
					<Checkbox {...props} bind:checked={$formData.useDocumentTypeColumn} />
					<div>
						<Form.Label>Använd dokumenttypkolumn?</Form.Label>
					</div>
				{/snippet}
			</Form.Control>
		</Form.Field>

		{#if $formData.useDocumentTypeColumn}
			<Form.Field {form} name="documentType">
				<Form.Control>
					{#snippet children({ props })}
						<Form.Label>Dokumenttypkolumn</Form.Label>
						<Input {...props} bind:value={$formData.documentTypeColumn} type="text" />
					{/snippet}
				</Form.Control>
				<Form.Description>Namn på kolumnen.</Form.Description>
				<Form.FieldErrors />
			</Form.Field>
		{:else}
			<Form.Field {form} name="documentType">
				<Form.Control>
					{#snippet children({ props })}
						<Form.Label>Dokumenttyp</Form.Label>
						<Select.Root type="single" bind:value={$formData.documentType} name={props.name}>
							<Select.Trigger {...props} class="w-full" />
							<Select.Content class="w-full">
								<Select.Item value="Produktdatablad" label="Produktdatablad" />
								<Select.Item value="Säkerhetsdatablad" label="Säkerhetsdatablad" />
								<Select.Item value="Prestandadeklaration" label="Prestandadeklaration" />
							</Select.Content>
						</Select.Root>
					{/snippet}
				</Form.Control>
				<Form.Description>Välj dokumenttypen du vill lägga till.</Form.Description>
				<Form.FieldErrors />
			</Form.Field>
		{/if}

		<Form.Field {form} name="sheetName">
			<Form.Control>
				{#snippet children({ props })}
					<Form.Label>Excel-ark</Form.Label>
					<Input {...props} bind:value={$formData.sheetName} type="text" />
				{/snippet}
			</Form.Control>
			<Form.Description>Namn på excel-arket i filen.</Form.Description>
			<Form.FieldErrors />
		</Form.Field>
	</div>

	<!-- Company -->
	<Form.Field {form} name="company">
		<Form.Control>
			{#snippet children({ props })}
				<Form.Label>Märke</Form.Label>
				<Select.Root type="single" bind:value={$formData.company} name={props.name}>
					<Select.Trigger {...props} class="w-full" />
					<Select.Content class="w-full">
						<Select.Item value="illbruck" label="Illbruck" />
						<Select.Item value="nullifire" label="Nullifire" />
						<Select.Item value="matacryl" label="Matacryl" />
						<Select.Item value="flowcrete" label="Flowcrete" />
						<Select.Item value="vandex" label="Vandex" />
						<Select.Item value="tremco" label="Tremco" />
					</Select.Content>
				</Select.Root>
			{/snippet}
		</Form.Control>
		<Form.FieldErrors />
	</Form.Field>

	<!-- File picker -->
	<Form.Field {form} name="fileName">
		<Form.Control>
			<Form.Label>Excel-fil</Form.Label>
			<div class="flex items-center gap-2">
				<Button type="button" onclick={selectFile}>Välj fil</Button>
				<span class="truncate text-sm italic">{selectedFilePath}</span>
			</div>
		</Form.Control>
		<Form.FieldErrors />
	</Form.Field>

	<!-- Column Config -->
	<div class="mt-6 space-y-4 rounded-lg border p-4">
		<h3 class="text-lg font-semibold">Kolumnkonfiguration</h3>

		<div class="grid grid-cols-1 gap-3 sm:grid-cols-2">
			<Form.Field {form} name="articleNumberColumn">
				<Form.Control>
					{#snippet children({ props })}
						<Form.Label>Artikelnummerkolumn</Form.Label>
						<Input {...props} bind:value={$formData.articleNumberColumn} type="text" />
					{/snippet}
				</Form.Control>
				<Form.FieldErrors />
			</Form.Field>

			<Form.Field {form} name="pdfLinkColumn">
				<Form.Control>
					{#snippet children({ props })}
						<Form.Label>PDF-länkkolumn</Form.Label>
						<Input {...props} bind:value={$formData.pdfLinkColumn} type="text" />
					{/snippet}
				</Form.Control>
				<Form.FieldErrors />
			</Form.Field>
		</div>

		<!-- <Form.Field {form} name="updateDate" class="flex gap-3 rounded-md border p-4">
			<Form.Control>
				{#snippet children({ props })}
					<Checkbox {...props} bind:checked={$formData.updateDate} />
					<div>
						<Form.Label>Uppdatera datum?</Form.Label>
						<Form.Description>Aktivera för att skriva dagens datum i Excel.</Form.Description>
					</div>
				{/snippet}
			</Form.Control>
		</Form.Field> -->

		<!-- {#if $formData.updateDate}
			<Form.Field {form} name="dateColumn">
				<Form.Control>
					{#snippet children({ props })}
						<Form.Label>Datumkolumn</Form.Label>
						<Input {...props} bind:value={$formData.dateColumn} type="text" />
					{/snippet}
				</Form.Control>
				<Form.FieldErrors />
			</Form.Field>
		{/if} -->
	</div>

	<!-- Submit Button -->
	<div class="flex justify-end">
		<Form.Button disabled={loading} class="mt-8 w-full sm:w-auto">
			{#if loading}
				<Loader class="h-4 w-4 animate-spin" /> Arbetar...
			{:else}
				Börja leta efter dokument
			{/if}
		</Form.Button>
	</div>

	<!-- Progress UI -->
	{#if cancelled}
		<p class="py-14 text-center text-lg font-medium">Operationen avbröts</p>
	{:else if progress}
		<div class="py-8">
			<Progress value={progress} />
			<div class="flex justify-between py-2">
				{#if progress < 100}
					<Button type="button" onclick={() => CancelFetch()}>Avbryt</Button>
				{/if}
				<p class="text-sm font-medium">{progress.toFixed(2)}%</p>
			</div>
		</div>
	{/if}
</form>
