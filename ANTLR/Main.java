// Multumesc echipei de CPL pentru acest schelet Main

import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.tree.*;

import java.io.IOException;
import java.util.ArrayList;
import java.io.*;

public class Main {
    public static int processNum = 0;
    public static ArrayList processFinish;
    public static void main(String[] args) throws IOException {

        HelloLexer lexer1 = null;
        HelloLexer lexer2 = null;
        CommonTokenStream tokenStream1 = null;
        CommonTokenStream tokenStream2 = null;
        HelloParser parser1 = null;
        HelloParser parser2 = null;
        ParserRuleContext globalTree = null;

        // True if any lexical or syntax errors occur.
        boolean lexicalSyntaxErrors = false;

        // Deschidem fisierul input pentru a incepe parsarea
        String fileName1 = "input1";
        CharStream input1 = CharStreams.fromFileName(fileName1);
        String fileName2 = "input2";
        CharStream input2 = CharStreams.fromFileName(fileName2);
        try {
            System.setOut(new PrintStream(new File("firstOrder_Output")));
        } catch (Exception e) {
            e.printStackTrace();
        }

        // Definim Lexer-ul
        lexer1 = new HelloLexer(input1);
        lexer2 = new HelloLexer(input2);
        // Obtinem tokenii din input
        tokenStream1 = new CommonTokenStream(lexer1);
        tokenStream2 = new CommonTokenStream(lexer2);
        // Definim Parser-ul
        parser1 = new HelloParser(tokenStream1);
        parser2 = new HelloParser(tokenStream2);
        // Incepem parsarea
        ParserRuleContext tree1 = parser1.mainnode();
        ParserRuleContext tree2 = parser2.mainnode();
        // Vizitam AST-ul
        System.out.println("pc=m∧pc'=&∧pc1'=La0∧pc2'=Lb0 ∨"); 
        processFinish = new ArrayList();
        MyVisitor visitor = new MyVisitor();
        visitor.visit(tree1);
        processNum++;
        visitor = new MyVisitor();
        visitor.visit(tree2);

        System.out.println("pc=&∧pc'=m'∧pc1="+processFinish.get(0)+"∧pc2="+ processFinish.get(0) +"∧pc1'=&∧pc2'=& ∨");
    }
}
