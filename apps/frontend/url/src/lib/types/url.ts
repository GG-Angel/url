import { z } from 'zod';

export const urlSchema = z.object({
	id: z.number().nonnegative(),
	url: z.url(),
	code: z.string(),
	created_at: z.coerce.date()
});

export type Url = z.infer<typeof urlSchema>;
