<template>
    <v-container class="grey lighten-5">
        <v-row no-gutters justify="center">
            <v-col v-for="p, key in sortedData" :key="key" sm="3" class="mb-6 mx-2" style="min-width: 180px;">
                <v-card class="pa-2" tile outlined style="text-align: center;">
                    <v-card-title>
                        {{ nameFormatter(key) }}
                    </v-card-title>
                    <v-progress-circular bg-color="transparent" :rotate="-90" :size="150" :width="25" :model-value="calcProgress(key, p)" :color="colorFormatter(settings, key, p)" style="margin-bottom: -50px;">
                        {{ p }} {{ unitFormatter(key, p) }}
                    </v-progress-circular>
                </v-card>
            </v-col>
        </v-row>
    </v-container>
</template>

<script>
import { nameFormatter, unitFormatter, colorFormatter } from "@/utils"

export default {
    name: 'controller-real-time',
    data() {
        return {
            tab: "controller_real_time"
        }
    },
    props: {
        data: Object,
        settings: Object,
    },
    methods: {
        calcProgress(name, value) {
            if (name.includes("temperature")) { return this._calcProgress(value, 0, 40); } 
            else if (name.includes("voltage")) { return this._calcProgress(value, 0, 16); }
            else if (name.includes("power")) { return this._calcProgress(value, 0, 120); }
            else if (name.includes("current")) { return this._calcProgress(value, 0, 10); }
            return this._calcProgress(value) 
        },
        _calcProgress(value, min, max) {
            if (min>=0 && max >= 0){
                const coef = 50 / (max - min);
                value = coef * Math.abs(value);
            } else {
                value = value <= 50 ? value : value <= 100 ? value/2 : value <= 1000 ? value/10/2 : 50
            }
            return value > 50 ? 50 : value;
        },
        nameFormatter,
        unitFormatter,
        colorFormatter
    },
    computed: {
        sortedData() {
            return this.data//?.sort()
        },
    }
 }
</script>
