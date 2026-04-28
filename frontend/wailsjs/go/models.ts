export namespace main {
	
	export class FetchParams {
	    file: string;
	    startNumber: number;
	    endNumber: number;
	    sheetName: string;
	    artikelnummerCol: string;
	    pdfLinkCol: string;
	    documentType: string;
	    use_document_type_column: boolean;
	    document_type_column: string;
	    company: string;
	    update_date: boolean;
	    date_column: string;
	
	    static createFrom(source: any = {}) {
	        return new FetchParams(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.file = source["file"];
	        this.startNumber = source["startNumber"];
	        this.endNumber = source["endNumber"];
	        this.sheetName = source["sheetName"];
	        this.artikelnummerCol = source["artikelnummerCol"];
	        this.pdfLinkCol = source["pdfLinkCol"];
	        this.documentType = source["documentType"];
	        this.use_document_type_column = source["use_document_type_column"];
	        this.document_type_column = source["document_type_column"];
	        this.company = source["company"];
	        this.update_date = source["update_date"];
	        this.date_column = source["date_column"];
	    }
	}

}

