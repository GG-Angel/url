import { URL_SERVICE_URL } from '$env/static/private';
import { error } from '@sveltejs/kit';
import type { PageServerLoad } from './$types';
import { getUrlsSchema } from '$lib/types/api';

export const load: PageServerLoad = async ({ fetch }) => {
	const res = await fetch(`${URL_SERVICE_URL}/urls`);

	if (!res.ok) {
		error(res.status, 'Failed to fetch URLs');
	}

	const urls = getUrlsSchema.parse(await res.json());
	return { urls };
};

export const actions = {
	create: async ({ request }: { request: Request }) => {
		const formData = await request.formData();
		const url = formData.get('url');

		await fetch(`${URL_SERVICE_URL}/urls`, {
			method: 'POST',
			body: JSON.stringify({ url }),
			headers: { 'Content-Type': 'application/json' }
		});

		return { success: true };
	}
};
