import { z } from 'zod';
import { urlSchema } from './url';

export const getUrlsSchema = z.array(urlSchema);

export const createUrlSchema = z.object({
	url: z.url()
});

export type GetUrlsResponse = z.infer<typeof getUrlsSchema>;
export type CreateUrlRequest = z.infer<typeof createUrlSchema>;
