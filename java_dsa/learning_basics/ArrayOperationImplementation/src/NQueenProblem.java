import java.util.Arrays;
import java.util.Scanner;

public class NQueenProblem {
    public static int N = 1;

    public static void main(String[] args) {
        System.out.println("Welcome to N Queen problem, Enter value of N");
        Scanner in = new Scanner(System.in);
        N = in.nextInt();
        if (N <= 0){
            System.out.println("N cannot be negative or zero");
            return;
        }
        int[][] board = new int[N][N];
        if (!solveNQueens(board, 0)) {
            System.out.println("Unable to find any solution");
        }
    }

    static boolean isSafe(int[][] board, int row, int column) {
        // present in the same row
        for (int x = 0; x < column; x++) {
            if (board[row][x] == 1) {
                return false;
            }
        }

        // upper left diagonal
        for (int x = row, y = column; x >= 0 && y >= 0; x--, y--) {
            if (board[x][y] == 1) {
                return false;
            }
        }

        // lower right diagonal
        for (int x = row, y = column; x < N && y >= 0; x++, y--) {
            if (board[x][y] == 1) {
                return false;
            }
        }
        return true;
    }

    static boolean solveNQueens(int[][] board, int column) {
        // solution found
        if (column == N) {
            for (int[] row : board) {
                System.out.println(Arrays.toString(row));
            }
            System.out.println();
            return true;
        }

        for (int i = 0; i < N; i++) {
            if (isSafe(board, i, column)) {
                board[i][column] = 1;
                if (solveNQueens(board, column + 1)) {
                    return true;
                }
                board[i][column] = 0;
            }
        }

        // no solution found
        return false;
    }
}
