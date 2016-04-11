class Main {
    public static int[] qsort(int[] lst) {
        // return Array(1, 2, 3);
        int[] myList = {1, 2, 3};
        return myList;
    }

    public static void main(String[] args) {
        // System.out.println("Yo");
        // System.out.println(qsort(new int[]{3, 5, 8, 10, 2, 4, 6, 1, 7, 8}));

        int[] inputArray = new int[]{3, 5, 8, 10, 2, 4, 6, 1, 7, 8};
        for (int x : inputArray) {
            System.out.println(x);
        }
        int[] myResult = qsort(inputArray);
        for (int x : myResult) {
            System.out.println(x);
        }
    }
}
