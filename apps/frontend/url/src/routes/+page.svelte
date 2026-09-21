<script lang="ts">
	import { resolve } from '$app/paths';
	import Button from '$lib/components/Button.svelte';
	import CardWithHeader from '$lib/components/CardWithHeader.svelte';
	import type { PageProps } from './$types';
	import CopyButton from '$lib/components/CopyButton.svelte';

	let { data }: PageProps = $props();
</script>

<CardWithHeader header="Create a short URL">
	<form class="flex flex-col gap-4" method="POST" action="?/create">
		<label class="sr-only" for="url">URL to shorten</label>
		<input name="url" type="url" placeholder="URL to be shortened" required />
		<div class="flex gap-4">
			<label class="sr-only" for="slug">Custom slug</label>
			<input name="slug" type="text" placeholder="Custom slug" />
			<label class="sr-only" for="slug">Tags</label>
			<input name="tags" type="text" placeholder="Add tags to the URL" />
		</div>
		<Button class="self-start">Save</Button>
	</form>
</CardWithHeader>

<CardWithHeader
	header="Recently created URLs"
	link={{ href: resolve('/urls'), text: 'See all' }}
>
	{#if data.urls.length === 0}
		<p
			class="flex min-h-32 shrink-0 items-center justify-center text-sm text-subtle-foreground"
		>
			No URLs yet.
		</p>
	{/if}

	{#if data.urls.length > 0}
		<div class="table">
			<div class="table-header-group">
				<div class="table-row font-semibold">
					<div class="table-cell">Created At</div>
					<div class="table-cell">Code</div>
					<div class="table-cell">URL</div>
					<div class="table-cell">Tags</div>
					<div class="table-cell">Visits</div>
				</div>
			</div>
			<div class="table-row-group">
				{#each data.urls as url (url.id)}
					<div class="table-row">
						<div class="table-cell">{url.created_at.toLocaleString()}</div>
						<div class="table-cell">
							<div class="flex items-center gap-3">
								{url.code}
								<CopyButton copyText={url.code} />
							</div>
						</div>
						<div class="table-cell">
							{url.url}
						</div>
						<div class="table-cell">
							{#if url.tags.length > 0}
								{#each url.tags as tag (tag.id)}
									<span>{tag.name}</span>
								{/each}
							{:else}
								<span>None</span>
							{/if}
						</div>
						<div class="table-cell">35</div>
					</div>
				{/each}
			</div>
		</div>
	{/if}
</CardWithHeader>
