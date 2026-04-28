<script lang="ts">
	import { page } from '$app/state';
	import { cn } from '@/utils';

	const links = [
		{
			name: 'Bilder',
			href: '/images'
		},
		{
			name: 'Dokument',
			href: '/documents'
		}
	];
</script>

<nav class="container mx-auto flex items-center gap-3 pt-12">
	{#each links as link}
		<a
			href={link.href}
			class={cn('', {
				'underline opacity-30': page.url.pathname.includes(link.href)
			})}>{link.name}</a
		>
	{/each}
</nav>
