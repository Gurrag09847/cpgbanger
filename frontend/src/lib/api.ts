import type {
	UploadResponse,
	UploadRow,
	JobRow,
	FetchRequest,
	FetchResponse,
	CancelRequest,
	Settings
} from './types';

const BASE = '/api';

async function post<T>(path: string, body?: unknown): Promise<T> {
	const res = await fetch(`${BASE}${path}`, {
		method: 'POST',
		headers: body instanceof FormData ? {} : { 'Content-Type': 'application/json' },
		body: body instanceof FormData ? body : JSON.stringify(body)
	});
	if (!res.ok) {
		const err = await res.json().catch(() => ({ error: res.statusText }));
		throw new Error(err.error || res.statusText);
	}
	return res.json();
}

async function get<T>(path: string): Promise<T> {
	const res = await fetch(`${BASE}${path}`);
	if (!res.ok) {
		const err = await res.json().catch(() => ({ error: res.statusText }));
		throw new Error(err.error || res.statusText);
	}
	return res.json();
}

async function del<T>(path: string): Promise<T> {
	const res = await fetch(`${BASE}${path}`, { method: 'DELETE' });
	if (!res.ok) {
		const err = await res.json().catch(() => ({ error: res.statusText }));
		throw new Error(err.error || res.statusText);
	}
	return res.json();
}

export function uploadFile(file: File): Promise<UploadResponse> {
	const fd = new FormData();
	fd.append('file', file);
	return post<UploadResponse>('/upload', fd);
}

export function listUploads(): Promise<UploadRow[]> {
	return get<UploadRow[]>('/uploads');
}

export function deleteUpload(id: string): Promise<{ status: string }> {
	return del<{ status: string }>(`/uploads/${id}`);
}

export function getUploadDownloadUrl(id: string): string {
	return `${BASE}/uploads/${id}/download`;
}

export function startFetch(req: FetchRequest): Promise<FetchResponse> {
	return post<FetchResponse>('/fetch', req);
}

export function cancelJob(req: CancelRequest): Promise<{ status: string }> {
	return post<{ status: string }>('/cancel', req);
}

export function listJobs(): Promise<JobRow[]> {
	return get<JobRow[]>('/jobs');
}

export function getJob(id: string): Promise<JobRow> {
	return get<JobRow>(`/jobs/${id}`);
}

export function getJobDownloadUrl(id: string): string {
	return `${BASE}/jobs/${id}/download`;
}

export function getJobSseUrl(id: string): string {
	return `${BASE}/jobs/${id}/sse`;
}

export function getExportUrl(): string {
	return `${BASE}/export`
}

export function getSettings(): Promise<Settings> {
	return get<Settings>('/settings');
}

export function updateSettings(settings: Record<string, string>): Promise<{ status: string }> {
	return post<{ status: string }>('/settings', settings);
}
