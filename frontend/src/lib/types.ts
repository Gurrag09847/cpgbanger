export interface UploadResponse {
	file_id: string;
	file_name: string;
	file_path: string;
}

export interface UploadRow {
	id: string;
	file_name: string;
	original_name: string;
	file_path: string;
	uploaded_at: string;
}

export interface JobRow {
	id: string;
	file_name: string;
	file_path: string;
	sheet_name: string;
	company: string;
	status: string;
	total_rows: number;
	processed_rows: number;
	progress: number;
	error_message: string;
	created_at: string;
	updated_at: string;
}

export interface JobUpdate {
	status: string;
	progress: number;
	processed: number;
	total: number;
	error: string;
}

export interface FetchRequest {
	upload_id?: string
	file_id?: string
	file_path?: string
	startNumber: number
	endNumber: number
	sheetName: string
	sheetNames?: string[]
	artikelnummerCol: string
	pdfLinkCol: string
	documentType: string
	use_document_type_column: boolean
	document_type_column: string
	company: string
	update_date: boolean
	date_column: string
	concurrency: number
}

export interface FetchResponse {
	job_id: string;
}

export interface CancelRequest {
	job_id: string;
}

export type Settings = Record<string, string>;
