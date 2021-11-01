package Eight;

import java.util.ArrayList;
import java.io.BufferedReader;
import java.io.File;
import java.io.FileReader;
import java.io.IOException;

class Statement{
    String action;
    String Same;
    String b_pc;
    String e_pc;

    int ib_pc;
    int ie_pc;
    Statement next;


    @Override
    public String toString() {
        return action+"\n"+Same+"\n"+b_pc+"\n"+e_pc;
    }


}

public class FOL_to_KS {


    public static void main(String[] args) throws IOException {
        ArrayList<Statement> state=new ArrayList<>();
        ArrayList<String> temp;
        temp=readFiles("F:\\Java_files\\javaSE_Code\\ieda_test\\src\\Eight\\data");
        for (String s : temp) {
            Statement test=spilt(s);
            System.out.println(test);
            //System.out.println(s);
        }
    }

    //获取一阶逻辑公式
    public static ArrayList<String> readFiles(String filename) throws IOException {
        File file =new File(filename);
        ArrayList<String> ret=new ArrayList<>();
        BufferedReader reader=null;
        reader=new BufferedReader(new FileReader(file));
        String tempString=null;
        while ((tempString= reader.readLine())!=null){
            //System.out.println(tempString);
            ret.add(tempString);
        }
        return ret;
    }

    //字符分割
    public static Statement spilt(String sentence){
        Statement ret=new Statement();
        int i=0;
        char[] d =sentence.toCharArray();
        String temp= "";
        for(char s : d){
            if(s=='∧'){
                //System.out.println(temp);
                i++;
                ret=add_info(ret,temp,i);
                temp = "";
            }
            else if(s=='∨'){
                //System.out.println(temp);
                i++;
                ret=add_info(ret,temp,i);
                temp = "";
            }
            else {
                if(s!=' '){
                    temp+=s;
                }
            }
        }
        return ret;
    }

    //信息转换
    public static Statement add_info(Statement state,String info,int index){
        if(index==1){
            state.action=info;
        }
        else if(index ==2){
            state.Same= info.substring(info.indexOf("(")+1,info.indexOf(")"));
        }
        else if(index==3){
            String[] temp=info.split("=");
            String[] temp1=temp[1].split("L");
            state.b_pc=temp1[1];
        }
        else{
            String[] temp=info.split("=");
            String[] temp1=temp[1].split("L");
            state.e_pc=temp1[1];
        }

        return state;
    }

    //利用action和same来判断状态
    public static void state_connect(ArrayList<Statement> state){

    }
}





