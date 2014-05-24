import com.sap.conn.jco.JCoDestination;
import com.sap.conn.jco.JCoException;
import com.sap.conn.jco.JCoDestinationManager;
import com.sap.conn.jco.JCoFunction;
import java.io.File;
import java.io.FileOutputStream;
import java.io.IOException;
import java.io.BufferedReader;
import java.io.FileReader;
import com.sap.conn.jco.AbapException;
import com.sap.conn.jco.JCoField;
import com.sap.conn.jco.JCoTable;
import com.sap.conn.jco.JCoRecordFieldIterator;
import com.sap.conn.jco.JCoStructure;
import java.util.*;

public class connectToAbap{
	private static String ABAP_AS = "ABAP_AS";

	private static void eshtablishing(){
		FileOutputStream fop = null;
		File file;
		try{
			file = new File("tempfile.txt");
			fop = new FileOutputStream(file);
			try{
				JCoDestination destination = JCoDestinationManager.getDestination(ABAP_AS);
/*				System.out.println("Attributes:");
        		System.out.println(destination.getAttributes());
        		System.out.println();*/
        		BufferedReader br = null;
        		try{
        			String sCurrentLine;
        			br = new BufferedReader(new FileReader("functionName.txt"));
        			sCurrentLine = br.readLine();
        			JCoFunction function = destination.getRepository().getFunction(sCurrentLine);
        			if(function==null)
            			throw new RuntimeException("RFC Enabled Function Module " + sCurrentLine + " not found in SAP.");
            		br = new BufferedReader(new FileReader("inputParams.txt"));
            		while ((sCurrentLine = br.readLine()) != null){
            			String parametername = sCurrentLine;
            			sCurrentLine = br.readLine();
            			String type = br.readLine();
            			if(type.equals("string"))
            				function.getImportParameterList().setValue(parametername, sCurrentLine);
            		}
            		try{
            			function.execute(destination);
        			}catch(AbapException e){
            			System.out.println("Error:" + e.getMessage());
            			return;
        			}
        			Iterator<JCoField> iterator = function.getExportParameterList().iterator();
        			String exportParams = "";
                    String structoutput = "";
                    String exportTableParams = "";
        			int newlineflag = 0;
        			while(iterator.hasNext()){
            			JCoField field = iterator.next();
            			if(!field.getTypeAsString().equals("STRUCTURE")){
            				exportParams = exportParams + field.getName() + "\n" + field.getTypeAsString() + "\n" + field.getValue().toString() + "\n";
            			}
            			else if(field.getTypeAsString().equals("STRUCTURE")){
            				JCoStructure exportStructure = function.getExportParameterList().getStructure(field.getName());
                            structoutput = structoutput + field.getName() + "\n";
                            for(int i = 0; i < exportStructure.getMetaData().getFieldCount(); i++){
                                structoutput = structoutput + exportStructure.getMetaData().getName(i) + "\n" + exportStructure.getString(i) + "\n";
                            }
                            structoutput = structoutput + "END;\n";
            			}
        			}
        			exportParams = exportParams + "END;\n";
                    if(function.getTableParameterList()!=null){
                        Iterator<JCoField> iteratorTable = function.getTableParameterList().iterator();
                        while(iteratorTable.hasNext()){
                            JCoField field = iteratorTable.next();
                            JCoTable codes = function.getTableParameterList().getTable(field.getName());
                            exportTableParams = exportTableParams + field.getName() + "\n";
                            JCoRecordFieldIterator iterator1 = codes.getRecordFieldIterator();
                            while(iterator1.hasNextField()){
                                JCoField fieldtab = iterator1.nextField();
                                exportTableParams = exportTableParams + fieldtab.getName() + "\n";
                            }
                            exportTableParams = exportTableParams + "ENDMETADATA;\n";
                            for (int i = 0; i < codes.getNumRows(); i++) 
                            {
                                codes.setRow(i);
                                JCoRecordFieldIterator iterator2 = codes.getRecordFieldIterator();
                                while(iterator2.hasNextField()){
                                    JCoField fieldval = iterator2.nextField();
                                    exportTableParams = exportTableParams + fieldval.getValue().toString() + "\n";
                                }
                            }
                            exportTableParams = exportTableParams + "END;\n";
                        }
                    }
        			File fileExport = new File("exportParameters.txt");
					FileOutputStream fopExport = new FileOutputStream(fileExport);
					try{
        				byte[] exportParameters = exportParams.getBytes();
        				fopExport.write(exportParameters);
						fopExport.flush();
						fopExport.close();
					}catch(IOException e){
        				System.out.println("Error from Java : " + e.getMessage());
        			}
                    File fileStructExport = new File("exportStructParameters.txt");
                    FileOutputStream fopStructExport = new FileOutputStream(fileStructExport);
                    try{
                        byte[] exportStructParameters = structoutput.getBytes();
                        fopStructExport.write(exportStructParameters);
                        fopStructExport.flush();
                        fopStructExport.close();
                    }catch(IOException e){
                        System.out.println("Error from Java : " + e.getMessage());
                    }
                    File fileTable = new File("exportTableParams.txt");
                    FileOutputStream foptabExport = new FileOutputStream(fileTable);
                    try{
                        byte[] exportTabParameters = exportTableParams.getBytes();
                        foptabExport.write(exportTabParameters);
                        foptabExport.flush();
                        foptabExport.close();
                    }catch(IOException e){
                        System.out.println("Error from Java : " + e.getMessage());
                    }
        		}catch(IOException e){
        			System.out.println("Error from Java : " + e.getMessage());
        		}
				String content = "success";
				byte[] contentInBytes = content.getBytes();
				fop.write(contentInBytes);
				fop.flush();
				fop.close();
        	}catch(JCoException e){
        		String content = e.getMessage();
				byte[] contentInBytes = content.getBytes();
				fop.write(contentInBytes);
				fop.flush();
				fop.close();
        	}
        }catch(IOException e){
        	System.out.println("Error from Java : " + e.getMessage());
        }
	}

	public static void main(String [] args){
		eshtablishing();
	}
}