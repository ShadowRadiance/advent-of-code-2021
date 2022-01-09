# frozen_string_literal: true

require "./pods/room"

module Pods
  RSpec.describe Room do
    subject { Room.new(door, letter: letter, capacity: capacity) }
    let(:door) { double }
    let(:letter) { "X" }
    let(:capacity) { 2 }
    let(:amphipod) { double(letter: "X") }
    
    describe "PUSHING AND POPPING" do
      context "with an empty room" do
        it "updates the amphipod when one is placed into the room" do
          expect(amphipod).to receive(:location=).with(subject)
          subject.push(amphipod)
        end
        
        it "raises an error when removing an amphipod" do
          expect { subject.pop }.to raise_error(Stack::Underflow)
        end
      end
      
      context "with one amphipod in the room" do
        let(:existing) { double(letter: "X") }
        before do
          allow(existing).to receive(:location=).with(subject)
          subject.push(existing)
        end

        it "updates the amphipod when one is placed into the room" do
          expect(amphipod).to receive(:location=).with(subject)
          subject.push(amphipod)
        end

        it "updates the amphipod when one is removed from the room" do
          expect(existing).to receive(:location=).with(nil)
          subject.pop
        end
      end

      context "with two (max) amphipods in the room" do
        let(:existing1) { double(letter: "X") }
        let(:existing2) { double(letter: "Y") }

        before do
          allow(existing1).to receive(:location=).with(subject)
          allow(existing2).to receive(:location=).with(subject)
          subject.push(existing1)
          subject.push(existing2)
        end

        it "can describe its contents" do
          expect(subject.amphipods.to_a).to eq([existing2, existing1])
        end

        it "updates the amphipods when removed from the room" do
          expect(existing1).to receive(:location=).with(nil)
          expect(existing2).to receive(:location=).with(nil)
          subject.pop
          subject.pop
        end
        
        it "raises an error when adding an amphipod" do
          expect(amphipod).to receive(:location=).with(subject).never
          
          expect { subject.push(amphipod) }.to raise_error(Stack::Overflow)
        end
      end
    end

    describe "#accessible_halls_with_cost" do
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

      context "with the example map step 0" do
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
          # consider the "C" at the top of the second room
          moves = roomB.accessible_halls_with_cost(100).sort_by {
            |move|
            move.to.index
          }

          expect(moves.map(&:from)).to all(eq(roomB))
          expect(moves.map(&:to)).to   eq([hall0,hall1,hall3,hall5,hall7,hall9,hall10])
          expect(moves.map(&:cost)).to eq([500,400,200,200,400,600,700])

        end
      end

      context "with the example map step 5" do
        let(:data) {
          <<~DATA
          #############
          #...........#
          ###D#B#C#D###
            #A#B#C#A#
            #########
          DATA
        }
        before {
          d_to_move = roomA.pop
          hall5.push(d_to_move)
          # => #############
          # => #.....D.....#
          # => ###.#B#C#D###
          # =>   #A#B#C#A#
          # =>   #########
        }

        it "determines the correct available moves" do
          # consider the "DA" in the last room
          moves = roomD.accessible_halls_with_cost(1000).sort_by {
            |move|
            move.to.index
          }

          expect(moves.map(&:from)).to eq([roomD,roomD,roomD ])
          expect(moves.map(&:to)).to   eq([hall7,hall9,hall10])
          expect(moves.map(&:cost)).to eq([2000, 2000, 3000  ])
        end
      end
    end
  end
end
  
  
  
  
  