<script lang="ts">
	import * as Form from '$lib/components/ui/form/index';
	import { Input } from '$lib/components/ui/input/index';
	import * as Select from '$lib/components/ui/select/index';
	import { formSchema, type FormSchema } from './schema';
	import { type SuperValidated, type Infer, superForm } from 'sveltekit-superforms';
	import { zod4Client } from 'sveltekit-superforms/adapters';
	import Loader from '@lucide/svelte/icons/loader';
	import { FetchDocuments, CancelFetch, SelectExcelFile } from '../../wailsjs/go/main/App';
	import * as rt from '../../wailsjs/runtime';
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
			cancelled = true; // Set the cancelled flag to true
			loading = false; // Optionally stop loading spinner
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
</form>
