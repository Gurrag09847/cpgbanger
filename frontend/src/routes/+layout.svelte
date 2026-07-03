<script lang="ts">
	import '../app.css';
	import { getExportUrl } from '../lib/api'
	let { children } = $props();
</script>

<div class="min-h-screen bg-background">
	<header class="border-border border-b bg-card">
		<div class="mx-auto flex max-w-6xl items-center justify-between px-6 py-4">
			<div>
				<h1 class="text-foreground text-xl font-bold">Tremco Dokumentationshämtare</h1>
				<p class="text-muted-foreground text-sm">Hämtar PDF-länkar för teknisk dokumentation</p>
			</div>
			<a
				href={getExportUrl()}
				class="border-border text-foreground hover:bg-secondary inline-flex items-center gap-2 rounded-md border px-3 py-2 text-sm font-medium transition-colors"
				download
			>
				<svg class="h-4 w-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
					<path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 16v1a3 3 0 003 3h10a3 3 0 003-3v-1m-4-4l-4 4m0 0l-4-4m4 4V4" />
				</svg>
				Exportera
			</a>
		</div>
	</header>
	<main class="mx-auto max-w-6xl px-6 py-8">
		{@render children()}
	</main>
</div>
