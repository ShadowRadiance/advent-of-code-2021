# frozen_string_literal: true

require "./pods/map"
require "./pods/move"

module Pods
  RSpec.describe Map do
    let(:map)    { Map.new(data) }
    let(:hall0)  { map.halls[0] }
    let(:hall1)  { map.halls[1] }
    let(:doorA)  { map.door("A") }
    let(:hall3)  { map.halls[2] }
    let(:doorB)  { map.door("B") }
    let(:hall5)  { map.halls[3] }
    let(:doorC)  { map.door("C") }
    let(:hall7)  { map.halls[4] }
    let(:doorD)  { map.door("D") }
    let(:hall9)  { map.halls[5] }
    let(:hall10) { map.halls[6] }
    let(:roomA)  { doorA.room }
    let(:roomB)  { doorB.room }
    let(:roomC)  { doorC.room }
    let(:roomD)  { doorD.room }
    let(:data) {
      <<~DATA
      #############
      #...........#
      ###B#C#B#D###
        #A#D#C#A#
        #########
      DATA
    }

    describe "#state" do
      context "with the starting example" do
        it "determines the current state" do
          expect(map.state).to eq(".|.|BA|.|CD|.|BC|.|DA|.|.")
        end
      end

      context "with a finished map" do
        let(:data) {
          <<~DATA
          #############
          #...........#
          ###A#B#C#D###
            #A#B#C#D#
            #########
          DATA
        }
        it "determines the current state" do
          expect(map.state).to eq(".|.|AA|.|BB|.|CC|.|DD|.|.")
        end
      end
    end
    
    describe "#restore_state" do
      let(:map) { Map.new(".|.|BA|.|B|D|CC|.|DA|.|.") }
      
      it "restores the state" do
        # EXPECTED STATE
        # #############
        # #.....D.....#
        # ###B#.#C#D###
        #   #A#B#C#A#
        #   #########
        expect(hall5.amphipod&.letter).to eq("D")
        map.halls.reject {|h| h.equal?(hall5) }.each do |hall|
          expect(hall.amphipod).to be_nil
        end
        expect(roomA.amphipods.map(&:letter).join).to eq("BA")
        expect(roomB.amphipods.map(&:letter).join).to eq("B")
        expect(roomC.amphipods.map(&:letter).join).to eq("CC")
        expect(roomD.amphipods.map(&:letter).join).to eq("DA")

        expect(map.state).to eq(".|.|BA|.|B|D|CC|.|DA|.|.")
      end
    end
    
    describe "#all_possible_moves" do
      context "with the starting example" do
        # #############
        # #...........#
        # ###B#C#B#D###
        #   #A#D#C#A#
        #   #########
        it "determines the correct moves" do
          moves = map.all_possible_moves
          expect(moves.size).to eq(28)
          expect(moves.map(&:to_a)).to contain_exactly(
            [roomA, hall0,  30],
            [roomA, hall1,  20],
            [roomA, hall3,  20],
            [roomA, hall5,  40],
            [roomA, hall7,  60],
            [roomA, hall9,  80],
            [roomA, hall10, 90],

            [roomB, hall0,  500],
            [roomB, hall1,  400],
            [roomB, hall3,  200],
            [roomB, hall5,  200],
            [roomB, hall7,  400],
            [roomB, hall9,  600],
            [roomB, hall10, 700],

            [roomC, hall0,  70],
            [roomC, hall1,  60],
            [roomC, hall3,  40],
            [roomC, hall5,  20],
            [roomC, hall7,  20],
            [roomC, hall9,  40],
            [roomC, hall10, 50],

            [roomD, hall0,  9000],
            [roomD, hall1,  8000],
            [roomD, hall3,  6000],
            [roomD, hall5,  4000],
            [roomD, hall7,  2000],
            [roomD, hall9,  2000],
            [roomD, hall10, 3000],
          )
        end
      end

      context "after 4 moves" do
        # let(:moves) {
        #   [
        #     Move.new(from: roomC, to: hall3, cost: 40),
        #     Move.new(from: roomB, to: roomC, cost: 400),
        #     Move.new(from: roomB, to: hall5, cost: 3000),
        #     Move.new(from: hall3, to: roomB, cost: 30),
        #   ]
        # }
        
        let(:restored) { Map.new(".|.|BA|.|B|D|CC|.|DA|.|.") }
        # EXPECTED STATE
        # #############
        # #.....D.....#
        # ###B#.#C#D###
        #   #A#B#C#A#
        #   #########
        
        it "determines the correct moves" do
          moves = restored.all_possible_moves
          expect(moves.size).to eq(7)
          expect(moves.map(&:to_s)).to contain_exactly(
            # the D in the hallway cannot move
            "[roomA->hall0]:30",
            "[roomA->hall1]:20",
            "[roomA->hall3]:20",
            # direct room-to-room move
            "[roomA->roomB]:40",
            # roomB has no moves (the one amphipod there is happy)
            # roomC has no moves (the two amphipods there are happy)
            "[roomD->hall7]:2000",
            "[roomD->hall9]:2000",
            "[roomD->hall10]:3000",
          )
        end
      end

      context "with a finished map" do
        let(:data) {
          <<~DATA
          #############
          #...........#
          ###A#B#C#D###
            #A#B#C#D#
            #########
          DATA
        }
        it "determines the correct moves" do
          expect(map.all_possible_moves.size).to eq(0)
        end
      end
    end

  end
end
