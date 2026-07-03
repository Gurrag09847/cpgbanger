import { writable } from 'svelte/store';

export type Toast = { id: string; message: string; variant: 'success' | 'error' | 'info' };

export const toasts = writable<Toast[]>([]);

const timeoutIds: Record<string, ReturnType<typeof setTimeout>> = {};

let currentToasts: Toast[] = [];
toasts.subscribe((v) => {
	currentToasts = v;
});

export function addToast(message: string, variant: Toast['variant'] = 'success') {
	const id = crypto.randomUUID();
	toasts.update((t) => [...t, { id, message, variant }]);
	timeoutIds[id] = setTimeout(() => {
		toasts.update((t) => t.filter((x) => x.id !== id));
		delete timeoutIds[id];
	}, 4000);
}

export function removeToast(id: string) {
	toasts.update((t) => t.filter((x) => x.id !== id));
	if (timeoutIds[id]) {
		clearTimeout(timeoutIds[id]);
		delete timeoutIds[id];
	}
}
