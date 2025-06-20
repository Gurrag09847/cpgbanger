export namespace main {
	
	export class FetchDocumentsParams {
	    file: string;
	    startNumber: number;
	    endNumber: number;
	    sheetName: string;
	    documentType: string;
	
	    static createFrom(source: any = {}) {
	        return new FetchDocumentsParams(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.file = source["file"];
	        this.startNumber = source["startNumber"];
	        this.endNumber = source["endNumber"];
	        this.sheetName = source["sheetName"];
	        this.documentType = source["documentType"];
	    }
	}

}

