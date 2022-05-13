<template>
    <v-container class="grey lighten-5">
        <v-row no-gutters justify="center">
            <v-col v-for="p, key in sortedData" :key="key" sm="3" class="mb-6 mx-2">
                <v-card class="pa-2" tile outlined style="text-align: center;">
                    <v-card-title>
                        {{ nameFormatter(key) }}
                    </v-card-title>
                    <v-progress-circular bg-color="transparent" :rotate="-90" :size="150" :width="30" :model-value="calcProgress(key, p)" color="primary" style="margin-bottom: -50px;">
                        {{ p }} {{ unit(key, p) }}
                    </v-progress-circular>
                </v-card>
            </v-col>
        </v-row>
    </v-container>
</template>

<script>

export default {
    name: 'controller-real-time',
    data() {
        return {
            tab: "controller_real_time"
        }
    },
    props: {
        data: Object,
        nameFormatter: Function,
        unit: Function
    },
    methods: {
        calcProgress(name, value) {
            switch (true) {
                case name.includes("temperature"): return this._calcProgress(0, 100, value);
                case name.includes("voltage"): return this._calcProgress(0, 16, value);
                case name.includes("power"): return this._calcProgress(0, 120, value);
                case name.includes("current"): return this._calcProgress(0, 10, value);
            }
            return value < 50 ? value : value < 100 ? value/2 : value < 1000 ? value/10/2 : 50 
        },
        _calcProgress(min, max, value) {
            const coef = 50 / (max - min);
            return coef * Math.abs(value);
        },
    },
    computed: {
        sortedData() {
            return this.data//?.sort()
        },
    }
 }
</script>
