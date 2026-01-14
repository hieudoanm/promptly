import { GeminiModel } from '@chat/clients/gemini/gemini.enums';
import { OpenRouterModel } from '@chat/clients/openrouter/openrouter.enums';
import { generate } from '@chat/services/genai.service';
import { summariseTranscript } from '@chat/services/youtube.service';
import { tryCatch } from '@chat/utils/try-catch';
import { z } from 'zod';
import { procedure, router } from '../trpc';

export const appRouter = router({
	genai: {
		generate: procedure
			.input(
				z.object({
					messages: z
						.object({
							role: z.enum(['ai', 'user']).default('user'),
							text: z.string().default(''),
						})
						.array()
						.default([]),
					model: z.string().default(''),
				}),
			)
			.mutation(async (options): Promise<{ text: string }> => {
				const { messages = [], model = GeminiModel.Gemini_2_5_Flash } =
					options.input;
				const { data, error } = await tryCatch(
					generate({ messages, model: model as GeminiModel | OpenRouterModel }),
				);
				if (error) {
					console.error('Error generating content:', error);
					return { text: 'An error occurred while generating content.' };
				}
				if (!data) {
					return { text: 'No response generated.' };
				}
				const output: string = data.output ?? 'No response generated.';
				return { text: output };
			}),
	},
	youtube: {
		transcript: {
			summarise: procedure
				.input(z.object({ videoId: z.string().default('') }))
				.mutation(async (options) => {
					const { videoId } = options.input;
					const { data, error } = await tryCatch(
						summariseTranscript({ videoId }),
					);
					if (error) {
						return { summary: error.message };
					}
					if (!data) {
						return { summary: 'No Summary.' };
					}
					return { summary: data.summary ?? 'No Summary.' };
				}),
		},
	},
});
// export type definition of API
export type AppRouter = typeof appRouter;
