import { error } from '@sveltejs/kit';
import type { RequestHandler } from './$types';

import { getSupabase } from '@supabase/auth-helpers-sveltekit';

import { stop } from '$lib/server/google/gmail';
import { getRefreshedGoogleAccessToken } from '$lib/supabase/client.server';

export const POST: RequestHandler = async (event) => {
	const { session, supabaseClient } = await getSupabase(event);
	if (!session) throw error(401, 'unauthorized');

	// get google refresh token
	let accessToken = '';
	try {
		accessToken = await getRefreshedGoogleAccessToken(supabaseClient);
	} catch (err) {
		// do something
		if (err instanceof Error) {
			throw error(500, err);
		}

		throw error(500, 'unexpected error occurred');
	}

	// stop watching for new emails
	const stopResponse = await stop(accessToken);
	if (stopResponse.status !== 200) throw error(500, 'failed to unsubscribe to gmail notifications');

	// "deactivate" the email in db
	const { error: updateError } = await supabaseClient
		.from('user_profile')
		.update({ is_active: false })
		.eq('user_id', session?.user.id);

	// should we re-subscribe if the db update fails?
	if (updateError) throw error(500, 'failed to saved changes to database');

	return new Response('success');
};
