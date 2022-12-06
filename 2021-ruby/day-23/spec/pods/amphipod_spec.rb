# frozen_string_literal: true

require "./pods/amphipod"

module Pods
  RSpec.describe Amphipod do
    describe "#valid_moves" do
      let(:map) { Map.new(data) }
      let(:hall0) { map.halls[0] }
      let(:hall1) { map.halls[1] }
      let(:doorA) { map.door("A") }
      let(:hall3) { map.halls[2] }
      let(:doorB) { map.door("B") }
      let(:hall5) { map.halls[3] }
      let(:doorC) { map.door("C") }
      let(:hall7) { map.halls[4] }
      let(:doorD) { map.door("D") }
      let(:hall9) { map.halls[5] }
      let(:hall10) { map.halls[6] }
      let(:roomA) { doorA.room }
      let(:roomB) { doorB.room }
      let(:roomC) { doorC.room }
      let(:roomD) { doorD.room }

      context "with the sample map at state 0" do
        let(:data) {
          <<~DATA
          #############
          #...........#
          ###B#C#B#D###
            #A#D#C#A#
            #########
          DATA
        }

        it "determines the correct available moves" do
          # each of the top amphipods should be able to move to any of the Hall locations
          all_the_halls = map.halls
          expect(roomA.amphipods[0].valid_moves.sort_by {|move| move.to.index}.map {|move| move.to }).to eq(all_the_halls)
          expect(roomB.amphipods[0].valid_moves.sort_by {|move| move.to.index}.map {|move| move.to }).to eq(all_the_halls)
          expect(roomC.amphipods[0].valid_moves.sort_by {|move| move.to.index}.map {|move| move.to }).to eq(all_the_halls)
          expect(roomD.amphipods[0].valid_moves.sort_by {|move| move.to.index}.map {|move| move.to }).to eq(all_the_halls)

          # check costs
          expect(roomA.amphipods[0].valid_moves.sort_by {|move| move.to.index}.map {|move| move.cost }).to eq([3, 2, 2, 4, 6, 8, 9].map{ |v| v * 10   })
          expect(roomB.amphipods[0].valid_moves.sort_by {|move| move.to.index}.map {|move| move.cost }).to eq([5, 4, 2, 2, 4, 6, 7].map{ |v| v * 100  })
          expect(roomC.amphipods[0].valid_moves.sort_by {|move| move.to.index}.map {|move| move.cost }).to eq([7, 6, 4, 2, 2, 4, 5].map{ |v| v * 10   })
          expect(roomD.amphipods[0].valid_moves.sort_by {|move| move.to.index}.map {|move| move.cost }).to eq([9, 8, 6, 4, 2, 2, 3].map{ |v| v * 1000 })

          # all the bottom amphipods should be stuck
          expect(roomA.amphipods[1].valid_moves).to be_empty
          expect(roomB.amphipods[1].valid_moves).to be_empty
          expect(roomC.amphipods[1].valid_moves).to be_empty
          expect(roomD.amphipods[1].valid_moves).to be_empty
        end
      end
      
      context "with the sample map at state 5" do
        let(:data) {
          <<~DATA
          #############
          #.....D.....#
          ###.#B#C#D###
            #A#B#C#A#
            #########
          DATA
        }

        it "determines the correct available moves" do
          # "D" in the corridor
          expect(hall5.amphipod.valid_moves).to be_empty
          
          # "A" in roomA (already correct)
          expect(roomA.amphipods[0].valid_moves).to be_empty

          # roomB (already correct)
          expect(roomB.amphipods[0].valid_moves).to be_empty
          expect(roomB.amphipods[1].valid_moves).to be_empty
          
          # roomC (already correct)
          expect(roomC.amphipods[0].valid_moves).to be_empty
          expect(roomC.amphipods[1].valid_moves).to be_empty
          
          # roomD
          expect(roomD.amphipods[0].valid_moves.sort_by {|move| move.to.index}.map {|move| move.to }).to eq([hall7, hall9, hall10])
          expect(roomD.amphipods[0].valid_moves.sort_by {|move| move.to.index}.map {|move| move.cost }).to eq([2,2,3].map{ |v| v * 1000 })
          expect(roomD.amphipods[1].valid_moves).to be_empty
        end
      end

      context "with another sample" do
        let(:data) {
          <<~DATA
          #############
          #.B...A.....#
          ###.#.#C#D###
            #A#B#C#D#
            #########
          DATA
        }

        it "determines the correct available moves" do
          # "B" in the corridor
          expect(hall1.amphipod.valid_moves.map {|move| move.to }).to eq([roomB])
          expect(hall1.amphipod.valid_moves.map {|move| move.cost }).to eq([4].map{ |v| v * 10   })
          
          # "A" in the corridor
          expect(hall5.amphipod.valid_moves.map {|move| move.to }).to eq([roomA])
          expect(hall5.amphipod.valid_moves.map {|move| move.cost }).to eq([4].map{ |v| v * 1    })
        end
      end

      context "disallow room-to-room with intervening amphipods" do
        let(:data) {
          <<~DATA
          #############
          #...B...D...#
          ###.#C#B#.###
            #A#D#C#A#
            #########
          DATA
        }
        it "determines the correct available moves" do
          # "B" in the corridor
          expect(hall3.amphipod.valid_moves).to be_empty
          # "D" in the corridor
          expect(hall7.amphipod.valid_moves).to be_empty
          # rooms
          expect(roomA.amphipods[0].valid_moves.map(&:to_s)).to be_empty # in correct place
          expect(roomB.amphipods[0].valid_moves.map(&:to_s)).to eq(["[roomB->hall5]:200"])
          expect(roomB.amphipods[1].valid_moves.map(&:to_s)).to be_empty
          expect(roomC.amphipods[0].valid_moves.map(&:to_s)).to eq(["[roomC->hall5]:20"])
          expect(roomC.amphipods[1].valid_moves.map(&:to_s)).to be_empty
          expect(roomD.amphipods[0].valid_moves.map(&:to_s)).to eq(["[roomD->hall9]:3", "[roomD->hall10]:4"])
        end
      end
    end
  end
end
