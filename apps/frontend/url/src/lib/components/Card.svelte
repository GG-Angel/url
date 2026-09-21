<script lang="ts">
	import type { ResolvedPathname } from "$app/types";
	import type { Snippet } from "svelte";

	type Props = {
        header: string;
        link?: {
            href: ResolvedPathname;
            text: string;
        }
		children: Snippet;
	};

    let { children, header, link }: Props = $props();
</script>

<section class="overflow-hidden rounded border bg-surface">
    <header class="flex items-center justify-between gap-4 border-b px-4 py-3 sm:px-5">
		<h2 class="font-medium">{header}</h2>
        {#if link}
            <a class="text-sm text-muted-foreground hover:text-foreground" href={link.href}>{link.text}</a>
        {/if}
	</header>

	<div class="flex flex-col gap-4 p-4" id="recent-urls">
		{@render children()}
	</div>
</section>