// @ts-check
import { defineConfig } from 'astro/config';
import mdx from '@astrojs/mdx';
import sitemap from '@astrojs/sitemap';

import playformCompress from '@playform/compress';

import purgecss from 'astro-purgecss';
import swup from '@swup/astro';

// https://astro.build/config
export default defineConfig({
    site: 'https://darwincereska.dev',
    integrations: [mdx(), sitemap(), purgecss(), playformCompress({
      CSS: true,
      HTML: {
        "html-minifier-terser": {
        removeAttributeQuotes: false,
        },
      },
      Image: true,
      JavaScript: true,
      SVG: true,
      }), swup({smoothScrolling: true, cache: true, preload: true, loadOnIdle: true, containers: ["main"], })],
});