import { z } from 'zod';
import { detailedTagSchema, urlWithTagsSchema } from './url';

export const getUrlsSchema = z.array(urlWithTagsSchema);

export const getTagsSchema = z.array(detailedTagSchema);

export const createUrlSchema = z.object({
	url: z.url()
});

export type GetUrlsResponse = z.infer<typeof getUrlsSchema>;
export type GetTagsResponse = z.infer<typeof getTagsSchema>;
export type CreateUrlRequest = z.infer<typeof createUrlSchema>;
