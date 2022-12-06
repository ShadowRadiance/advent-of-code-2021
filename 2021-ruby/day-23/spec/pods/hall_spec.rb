# frozen_string_literal: true

require "./pods/hall"

module Pods
  RSpec.describe Hall do
    subject { Hall.new(map, index) }

    let(:map) { double }
    let(:index) { 0 }
    let(:amphipod) { double }

    context "with an empty hall" do
      it "updates the amphipod when one is placed into the hall" do
        expect(amphipod).to receive(:location=).with(subject)
        subject.push(amphipod)
      end

      it "raises an error when removing an amphipod" do
        expect { subject.pop }.to raise_error(Stack::Underflow)
      end
    end

    context "with an amphipod in the hall" do
      before do
        allow(amphipod).to receive(:location=).with(subject)
        subject.push(amphipod)
      end
  
      it "updates the amphipod when one is removed from the hall" do
        expect(amphipod).to receive(:location=).with(nil)
        subject.pop
      end

      it "raises an error when adding an amphipod" do
        expect(amphipod).to receive(:location=).with(subject).never

        expect { subject.push(amphipod) }.to raise_error(Stack::Overflow)
      end
    end


  end
end

