import { defineCollection, z } from 'astro:content';

const blogCollection = defineCollection({type: 'content',schema: z.object({title: z.string(),description: z.string(),publishDate: z.coerce.date(),tags: z.array(z.string()),img: z.string(),img_alt: z.string().optional(),}),});
const workCollection = defineCollection({type: 'content',schema: z.object({title: z.string(),description: z.string(),publishDate: z.coerce.date(),tags: z.array(z.string()),img: z.string(),img_alt: z.string().optional(),}),});


export const collections = {
  'blog': blogCollection,
  'work': workCollection,
};
