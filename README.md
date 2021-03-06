#Algorithm


### 1. sort

    1-1. selection sort(선택정렬): 가장 작은것을 선택해서 제일 앞으로 보내는 정렬. O(N^2)
    1-2. bubble sort(버블정렬): 옆에 있는 값과 비교하여 더 작은값을 반복적으로 앞으로 보내는 정렬. O(N^2)
    1-3. insertion sort(삽입정렬): 자료 배열의 모든 요소를 앞에서부터 차례로 이미 정렬된 배열 부분과 비교하여, 자신의 위치를 찾아서 삽입하는 정렬. O(N^2)// 이미 정렬된 경우 가장 빠름.
    1-4. quick sort(퀵정렬): 특정한 값(피벗)을 기준으로 큰 숫자들과 작은 숫자들을 나눈 뒤 나눈 배열들을 다시 반복해서 정렬. O(N*logN), 최악 O(N^2) -> 이미 정렬된 경우.
    1-5. merge sort(병합정렬): 배열을 반으로 분할한 뒤 각자 계산하고 나중에 합치는 정렬. O(N*logN)
    1-6. heap sort(힙정렬): 힙트리 구조를 이용하는 정렬방법. O(N*logN)
        1) 이진 트리(binary tree): 모든 노드의 자식 노드가 2개이하인 트리구조.
        2) 힙: 최솟값이나 최댓값을 빠르게 찾아내기 위해 완전 이진트리를 기반으로 하는 트리.
    1-7: counting sort(계수정렬): 범위조건이 있는 경우 크기를 기준으로 갯수만 세서 정렬. O(N)

### 2. search
    2-1. BFS(breadth first search, 너비 우선 탐색)
        - 너비를 우선으로 탐색(1부터 가까운 노드들을 탐색). 최단 길이를 보장해야 할 때 사용(최단경로를 찾아줌). Queue사용.