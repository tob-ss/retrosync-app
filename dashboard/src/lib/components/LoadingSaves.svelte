<script lang="ts">
   
</script>

<div class="h-64 w-xl m-12 grid grid-cols-3 grid-rows-1 gap-4 p-12">
    <div class="col-span-2 col-start-2 row-start-1 grid grid-cols-3 grid-rows-3 gap-4 font-bold">
        <div class="col-span-3">
            <p class="skeleton rounded-xl">Value for Item 1</p>
        </div>
        <div class="col-span-3 row-start-2">
            <p class="skeleton rounded-xl">Value for Item 1</p>
        </div>
        <div class="row-start-3">
            <p class="skeleton rounded-xl">Value for Item 1</p>
        </div>
        <div class="col-start-3 row-start-3">
            <p class="skeleton rounded-xl">Value for Item 1</p>
        </div>
        
        
        
    </div>
    <div class="col-start-1 row-start-1 h-full w-full rounded-xl"> 
        <div class="skeleton h-full rounded-xl">Value for Item 1</div>
    </div>
</div>


<style>
.infoColumns {
    display: flex;
    flex-direction: row;
    flex-wrap: wrap;
    max-width: 650px;
    margin-top: 30px;
    margin-bottom: 16px;
}

.skeleton {
    animation: skeleton-loading 1s linear infinite alternate;
    color: transparent;
}

@keyframes skeleton-loading {
    0% {
        background-color: #2C2F48; /* FROM Color 1 */
    }
    100% {
        background-color: #E2E2E2; /* TO Color 2 */
    }
}
</style>