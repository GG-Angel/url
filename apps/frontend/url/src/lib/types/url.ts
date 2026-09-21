import { z } from 'zod';

export const tagSchema = z.object({
	id: z.number().nonnegative(),
	name: z.string(),
	created_at: z.coerce.date()
});

export const detailedTagSchema = tagSchema.extend({
	total_links: z.number().nonnegative()
});

export const urlSchema = z.object({
	id: z.number().nonnegative(),
	url: z.url(),
	code: z.string(),
	created_at: z.coerce.date()
});

export const urlWithTagsSchema = urlSchema.extend({
	tags: z.array(tagSchema)
});

export type Url = z.infer<typeof urlSchema>;
export type Tag = z.infer<typeof tagSchema>;
export type UrlWithTags = z.infer<typeof urlWithTagsSchema>;
export type DetailedTag = z.infer<typeof detailedTagSchema>;
