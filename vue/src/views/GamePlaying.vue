<template>
    <div class='container'>
        <div class='row my-3 p-3'>
            <div class='col'>
                Game ID: 12
                <br>
                Betting time: 13s
            </div>
            <div class='col text-center'>
            <button type="button" class="btn btn-danger">Quit Room</button>
            </div>
            <div class='col'>
                Username: TestUser
                <br>
                Token: 500
            </div>
        </div>
        <div class='divider'></div>
        <div class='row my-3'>
            <div class='col'>
                <Card :rank='r' :suit='s'></Card>
            </div>
            <div class='col'>
                <Card :rank='r' :suit='s'></Card>
            </div>
            <div class='col'>
                <Card v-for='i in 10' v-bind:key="i" :style="'position:absolute; left:' + (15 + i*1.5) + 'px; z-index: ' + i"></Card>
            </div>
            <div class='col'>
                <Card :rank='r' :suit='s'></Card>
            </div>
            <div class='col'>
                <Card :rank='r' :suit='s'></Card>
            </div>
        </div>
        <div class='row my-3'>
            <div class='col-5'>
                <UserList/>
            </div>
            <div class='col-3'>
                Amount
            </div>
            <div class='col-2'>
            </div>
            <div class='col-2'>
            </div>
        </div>
        <p>Test Animation</p>
        <button @click='r=""; s=""'>close card</button>
        <button @click='r="1"; s="S"'>open card</button>
        <button type="button" @click='showModal'>Launch demo modal</button>

        <!-- Modal -->
        <div class="modal fade" id="modal" tabindex="-1">
        <div class="modal-dialog">
            <div class="modal-content">
            <div class="modal-header">
                <h5 class="modal-title">Result</h5>
            </div>
            <div class="modal-body">
                <p>Winner: Banker</p>
                <p>Your bet: Banker</p>
                <p>Amount: 100</p>
            </div>
            <div class="modal-footer">
                <button type="button" class="btn btn-secondary" data-dismiss="modal">Close</button>
            </div>
            </div>
        </div>
        </div>
    </div>
</template>
 
<script>
import Card from "@/components/Card.vue";
import UserList from "@/components/UserList.vue";

export default{
    components: {
        Card,
        UserList
    },
    data: function() {
        return {
            r: '1',
            s: 'S'
        }
    },
    methods: {
        showModal() {
            $("#modal").modal('toggle');
        }
    }
}


</script> 

<style scoped>

</style>