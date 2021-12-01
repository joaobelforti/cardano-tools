#include <omp.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>


int main (int argc, char *argv[]) {
    
    int n_threads = 0;
    
    sscanf(argv[1], "%d", &n_threads);
    
    int n = 0;
    
    omp_set_num_threads(n_threads);

    #pragma omp parallel for

    for(n=0;n<n_threads;n++){
        
        char inputs[100]="";
        char snum[2]="";
        
        sprintf(snum, "%d", n+1);

        char input[100]="node sendTx.js";
        strcat(inputs," ");
        strcat(inputs,snum);
        strcat(inputs," ");
        strcat(inputs,argv[2]);
        strcat(inputs," ");
        strcat(inputs,argv[3]);
        
        strcat(input,inputs);
        system(input);
    }
}

//gcc botRun.c -o botRun -fopenmp
// ./botRun nThreads txValue addr1zzzzzzz
