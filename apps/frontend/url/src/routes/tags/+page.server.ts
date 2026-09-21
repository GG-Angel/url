import { URL_SERVICE_URL } from '$env/static/private';
import { error } from '@sveltejs/kit';
import type { PageServerLoad } from './$types';
import { getTagsSchema } from '$lib/types/api';

export const load: PageServerLoad = async ({ fetch }) => {
	const res = await fetch(`${URL_SERVICE_URL}/tags`);

	if (!res.ok) {
		error(res.status, 'Failed to fetch tags');
	}

	const tags = getTagsSchema.parse(await res.json());
	return { tags };
};
