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

<section>
    <header class="flex items-center justify-between gap-4">
		<h2>{header}</h2>
        {#if link}
            <a href={link.href}>{link.text}</a>
        {/if}
	</header>

	<div class="flex flex-col gap-4 p-4" id="recent-urls">
		{@render children()}
	</div>
</section>