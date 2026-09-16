<section>
    <header>
        <h1>Create a short URL</h1>
    </header>
    <form class="flex flex-col gap-4 p-4 sm:p-5" method="POST" action="?/create">
        <label class="sr-only" for="url">URL to shorten</label>
        <input
            id="url"
            name="url"
            placeholder="URL to be shortened"
            type="url"
            required
        />
        <div class="grid gap-4 sm:grid-cols-2">
            <label class="sr-only" for="slug">Custom slug</label>
            <input
                id="slug"
                name="slug"
                placeholder="Custom slug"
                type="text"
            />
            <label class="sr-only" for="tags">Tags</label>
            <input
                id="tags"
                name="tags"
                placeholder="Add tags to the URL"
                type="text"
            />
        </div>
        <button class="self-start rounded font-medium" type="submit">Save</button>
    </form>
</section>

<section>
    <header class="flex items-center justify-between gap-4">
        <h2>Recently created URLs</h2>
        <a href="#recent-urls">See all</a>
    </header>
    <div class="flex min-h-32 items-center justify-center px-4 py-8" id="recent-urls">
        <p class="text-subtle-foreground text-sm">No URLs yet.</p>
    </div>
</section>
